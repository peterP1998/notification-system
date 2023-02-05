package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaConsumer *kafka.Consumer

func CreateSubscriber() {
	var err error

	fmt.Printf("Starting consumer...")
	kafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id": "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics([]string{"slack-notification-topic"}, nil)

	go consumeMessages()
}

func consumeMessages() {

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			data := string(msg.Value)
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, data)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}