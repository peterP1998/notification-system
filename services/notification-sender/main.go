package main

import (
	"github.com/peterP1998/notification-system/notification-sender/config"
	"github.com/peterP1998/notification-system/notification-sender/consumer"
	"github.com/peterP1998/notification-system/notification-sender/consumer/retry"
	"github.com/peterP1998/notification-system/notification-sender/server"
	"github.com/peterP1998/notification-system/notification-sender/service"
	"log"
)

func main() {
	var configuration config.Config
	err := config.Read(&configuration)
	if err != nil {
		log.Fatal(err.Error())
	}
	consumer.CreateConsumers(configuration.KafkaHost, configuration.Topics, service.SenderServiceFacade{})
	retry.CreateRetryProducer(configuration.KafkaHost)
	server.Init(configuration.Host)
}
