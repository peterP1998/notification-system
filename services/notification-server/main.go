package main

import (
	"github.com/peterP1998/notification-system/notification-server/config"
	"github.com/peterP1998/notification-system/notification-server/server"
	"log"
)

func main() {
	var configuration config.Config
	err := config.Read(&configuration)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Starting application...")
	server.Init(configuration.Host, configuration.KafkaHost)
}
