package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/notification-sender/service"
)

var kafkaConsumer *kafka.Consumer

func CreateSubscriber() {
	var err error

	fmt.Printf("Starting consumer...")
	kafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics([]string{"slack-notification-topic", "email-notification-topic"}, nil)

	go consumeMessages()
}

func consumeMessages() {

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: \n", msg.Value)
			service.SendNotification(msg.Value)
			//fmt.Printf("Message on %s: \n", msg.Value)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
