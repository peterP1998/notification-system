package producer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"log"
)

var kafkaProducer *kafka.Producer

func InitProducer(kafkaHost string) {
	kafkaProducer, _ = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaHost})
	log.Printf("producer %v", kafkaProducer)

	go monitorEvents()
}

func ProduceMessage(notification *model.Notification, topic string) {
	b, _ := json.Marshal(notification)
	log.Printf("producer %v", kafkaProducer)
	err := kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func monitorEvents() {
	for e := range kafkaProducer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}
