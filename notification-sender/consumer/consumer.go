package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumeMessages() {
	fmt.Printf("Starting consumer...")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id": "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"slack-notification-topic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			data := string(msg.Value)
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, data)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}