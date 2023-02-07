package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/libs/notification/producer"
)

var kafkaProducer *kafka.Producer

func CreateProducer(kafkaHost string) {
	kafkaProducer, _ = producer.CreateProducer(kafkaHost)
}

func PublishNotification(notification []byte, topic string) {
	producer.ProduceMessage(kafkaProducer, notification, topic)
}
