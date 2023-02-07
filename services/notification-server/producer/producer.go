package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/libs/notification/producer"
	"log"
)

var kafkaProducer *kafka.Producer

type ProducerInterface interface {
	PublishNotification(notification []byte, topic string)
}

type Producer struct {
    kafkaProducer *kafka.Producer
}

func CreateProducer(kafkaHost string) ProducerInterface {
	log.Printf("Creating Producer")
	producer, _ := producer.CreateProducer(kafkaHost)
	return &Producer{
		kafkaProducer: producer,
	}
}

func (p Producer) PublishNotification(notification []byte, topic string) {
	log.Printf("Publishing Notification to topic %s", topic)
	producer.ProduceMessage(p.kafkaProducer, notification, topic)
}
