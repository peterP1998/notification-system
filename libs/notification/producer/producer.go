package producer

func CreateProducer(producer *kafka.Producer, kafkaHost string) {
	log.Println("Creating producer")
	producer, _ = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaHost})

	go monitorEvents(producer)
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
