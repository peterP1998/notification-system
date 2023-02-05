package producer

import (
	"log"
	"fmt"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/libs/notification/model"
)

var kafkaProducer *kafka.Producer

func InitProducer() { 
	fmt.Printf("Starting producer...")
	kafkaProducer, _ = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	log.Printf("producer %v", kafkaProducer)

	go func() {
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
	}()
}

func ProduceMessage(notification *model.Notification, topic string) {
	//data := []byte(notification)
	b, _ := json.Marshal(notification)
	log.Printf("producer %v", kafkaProducer)
	fmt.Println(topic)
	fmt.Println(b)
	err := kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
	
}