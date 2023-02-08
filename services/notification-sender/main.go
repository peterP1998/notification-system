package main

import (
	"github.com/peterP1998/notification-system/notification-sender/config"
	"github.com/peterP1998/notification-system/notification-sender/consumer"
	"github.com/peterP1998/notification-system/notification-sender/consumer/retry"
	"github.com/peterP1998/notification-system/notification-sender/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	var configuration config.Config
	err := config.Read(&configuration)
	log.Print(viper.Get("Email.From"))
	if err != nil {
		log.Fatal(err.Error())
	}
	consumer.CreateMainConsumer(configuration.KafkaHost, configuration.Topics)
	consumer.CreateRetryConsumer(configuration.KafkaHost)
	retry.CreateRetryProducer(configuration.KafkaHost)
	server.Init(configuration.Host)
}
