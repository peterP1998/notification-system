package main

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-server/config"
	"github.com/peterP1998/notification-system/notification-server/producer"
	"github.com/peterP1998/notification-system/notification-server/server"
	"log"
)

func main() {
	var configuration config.Config
	err := config.Read(&configuration)
	if err != nil {
		log.Fatal(err.Error())
	}
	producer.CreateProducer(configuration.KafkaHost)
	fmt.Println("Hello, World")
	server.Init(configuration.Host)
}
