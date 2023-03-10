package retry

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/libs/notification/producer"
	"log"
	"strconv"
	"strings"
)

var retryProducer *kafka.Producer

const (
	MAX_NUMBER_OF_RETRY = 5
	TOPIC_PREFIX        = "retry-topic-"
	DEAD_LETER_TOPIC    = "dead_letter_queue"
)

func CreateRetryProducer(kafkaHost string) {
	retryProducer, _ = producer.CreateProducer(kafkaHost)
}

func RetryMessage(msg *kafka.Message) {
	topic := ""
	log.Print("Retrying message wtih topic %s", msg.TopicPartition.Topic)
	if isAlreadyRetried(*msg.TopicPartition.Topic) {
		currentRetry := strings.TrimPrefix(*msg.TopicPartition.Topic, TOPIC_PREFIX)
		numberOfRetries, _ := strconv.Atoi(currentRetry)
		numberOfRetries += 1
		if numberOfRetries > MAX_NUMBER_OF_RETRY {
			log.Print("Adding message to dead letter queue")
			topic = DEAD_LETER_TOPIC
		} else {
			topic = TOPIC_PREFIX + strconv.Itoa(numberOfRetries)
		}
	} else {
		topic = TOPIC_PREFIX + "1"
	}
	producer.ProduceMessage(retryProducer, msg.Value, topic)
}

func GetRetryTopics() []string {
	var retryTopics [MAX_NUMBER_OF_RETRY]string
	for i := 1; i <= MAX_NUMBER_OF_RETRY; i++ {
		retryTopics[i-1] = TOPIC_PREFIX + strconv.Itoa(i)
	}
	return retryTopics[:]
}

func isAlreadyRetried(topic string) bool {
	if strings.Contains(topic, TOPIC_PREFIX) {
		return true
	}
	return false
}
