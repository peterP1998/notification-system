package main

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-sender/consumer"
	"github.com/peterP1998/notification-system/notification-sender/server"
	"github.com/peterP1998/notification-system/notification-sender/config"
	"log"
)

func main() {
	err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}
	consumer.CreateSubscriber()
	fmt.Println("Hello, World")
	server.Init()
}
