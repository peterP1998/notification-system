package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/notification-sender/service"
	"github.com/peterP1998/notification-system/notification-sender/config"
)

var kafkaConsumer *kafka.Consumer

func CreateSubscriber() {
	var err error

	fmt.Printf("Starting consumer...")
	kafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Configuration.KafkaHost,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(config.Configuration.Topics, nil)

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
