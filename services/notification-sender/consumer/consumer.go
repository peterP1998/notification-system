package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/notification-sender/service"
	"github.com/peterP1998/notification-system/notification-sender/consumer/retry"
)

const (
	MAX_NUMBER_OF_RETRY = 5
	TOPIC_PREFIX        = "retry_topic"
	DEAD_LETER_TOPIC    = "dead_letter_queue"
)

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

	fmt.Printf("Starting consumer...")
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
			fmt.Printf("Message on %s: \n", msg.Value)
			err = service.SendNotification(msg.Value)
			if err != nil {
				retry.RetryMessage(msg)
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
