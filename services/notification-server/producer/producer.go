package producer

import (
	"log"
	"fmt"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/peterP1998/notification-system/notification-server/config"
)

var kafkaProducer *kafka.Producer

func InitProducer() { 
	kafkaProducer, _ = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.Configuration.KafkaHost})
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