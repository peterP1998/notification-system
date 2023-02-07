package retry

import (
	"fmt"
	"github.com/peterP1998/notification-system/libs/notification/producer"
	"net/smtp"
	"strconv"
	"strings"
)

var retryProducer *kafka.Producer

const (
	MAX_NUMBER_OF_RETRY = 5
	TOPIC_PREFIX        = "retry_topic_"
	DEAD_LETER_TOPIC    = "dead_letter_queue"
)

func CreateRetryProducer(kafkaHost string) {
	producer.CreateProducer(retryProducer, kafkaHost)
}

func RetryMessage(msg *kafka.Message) {
	topic = ""
	if isAlreadyRetried(msg.TopicPartition.Topic) {
		current_retry = strings.TrimPrefix(msg.TopicPartition.Topic, TOPIC_PREFIX)
		number_of_retries = strconv.ParseInt(currentRetry) + 1
		if number_of_retries > MAX_NUMBER_OF_RETRY {
			topic = DEAD_LETER_TOPIC
		} else {
			topic = TOPIC_PREFIX + number_of_retries
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
