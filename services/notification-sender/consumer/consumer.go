package consumer

import (
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/notification-sender/consumer/retry"
	"github.com/peterP1998/notification-system/notification-sender/service"
)

var RETRY_TOPICS = []string{"retry-topic-1", "retry-topic-2", "retry-topic-3", "retry-topic-4", "retry-topic-5"}

var retryConsumer *kafka.Consumer
var kafkaConsumer *kafka.Consumer

func CreateRetryConsumer(kafkaHost string) {
	createConsumer(retryConsumer, kafkaHost, RETRY_TOPICS)
}

func CreateMainConsumer(kafkaHost string, kafkaTopics []string) {
	createConsumer(kafkaConsumer, kafkaHost, kafkaTopics)
}

func createConsumer(consumer *kafka.Consumer, kafkaHost string, kafkaTopics []string) {
	var err error

	log.Printf("Starting consumer for topics %v...", kafkaTopics)
	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaHost,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics(kafkaTopics, nil)

	go consumeMessages(consumer)
}

func consumeMessages(consumer *kafka.Consumer) {

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: \n", msg.Value)
			err = service.SendNotification(msg.Value)
			if err != nil {
				retry.RetryMessage(msg)
				log.Print(err)
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
