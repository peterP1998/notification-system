package main

import (
	"github.com/peterP1998/notification-system/notification-sender/config"
	"github.com/peterP1998/notification-system/notification-sender/consumer"
	"github.com/peterP1998/notification-system/notification-sender/server"
	"log"
)

func main() {
	var configuration config.Config
	err := config.Read(&configuration)
	if err != nil {
		log.Fatal(err.Error())
	}
	consumer.CreateSubscriber(configuration.KafkaHost, configuration.Topics)
	consumer.CreateRetryConsumer(configuration.KafkaHost)
	retry.CreateRetryProducer(configuration.KafkaHost)
	server.Init(configuration.Host)
}
