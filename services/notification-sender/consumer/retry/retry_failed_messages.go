package retry

import (
	"github.com/peterP1998/notification-system/libs/notification/producer"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	if isAlreadyRetried(*msg.TopicPartition.Topic) {
		currentRetry := strings.TrimPrefix(*msg.TopicPartition.Topic, TOPIC_PREFIX)
		numberOfRetries, _ := strconv.Atoi(currentRetry)
		numberOfRetries += 1
		if numberOfRetries > MAX_NUMBER_OF_RETRY {
			topic = DEAD_LETER_TOPIC
		} else {
			topic = TOPIC_PREFIX + strconv.Itoa(numberOfRetries)
		}
	} else {
		topic = TOPIC_PREFIX + "1"
	}
	producer.ProduceMessage(retryProducer, msg.Value, topic)
}

func isAlreadyRetried(topic string) bool {
	if strings.Contains(topic, TOPIC_PREFIX) {
		return true
	}
	return false
}
