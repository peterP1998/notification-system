package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/notification-sender/service"
)

var kafkaConsumer *kafka.Consumer

func CreateSubscriber(kafkaHost string, kafkaTopics []string) {
	var err error

	fmt.Printf("Starting consumer...")
	kafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaHost,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(kafkaTopics, nil)

	go consumeMessages()
}

func consumeMessages() {

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: \n", msg.Value)
			err = service.SendNotification(msg.Value)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
