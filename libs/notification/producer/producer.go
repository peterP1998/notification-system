package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func CreateProducer(kafkaHost string) (*kafka.Producer, error) {
	log.Println("Creating producer")
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaHost})
	if err != nil {
		return nil, err
	}

	go monitorEvents(producer)

	return producer, nil
}

func ProduceMessage(producer *kafka.Producer, notification []byte, topic string) {
	log.Printf("producer %v", producer)
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          notification,
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func monitorEvents(producer *kafka.Producer) {
	for e := range producer.Events() {
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
