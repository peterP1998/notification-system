package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/notification-sender/consumer/retry"
	"github.com/peterP1998/notification-system/notification-sender/service"
	"log"
)

var RETRY_TOPICS = []string{"retry-topic-1", "retry-topic-2", "retry-topic-3", "retry-topic-4", "retry-topic-5"}

var retryConsumer *kafka.Consumer
var kafkaConsumer *kafka.Consumer
var serviceFacade service.SenderServiceFacadeInterface

func CreateConsumers(kafkaHost string, kafkaTopics []string, serviceFacade service.SenderServiceFacadeInterface) {

	log.Print("Creating the main consumer")
	createConsumer(kafkaConsumer, kafkaHost, kafkaTopics)

	log.Print("Creating the retry consumer")
	createConsumer(kafkaConsumer, kafkaHost, kafkaTopics)

	serviceFacade = serviceFacade
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
			err = serviceFacade.SendNotification(msg.Value)
			if err != nil {
				retry.RetryMessage(msg)
				log.Print(err)
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
