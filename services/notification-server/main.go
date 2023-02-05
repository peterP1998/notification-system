package main

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-server/config"
	"github.com/peterP1998/notification-system/notification-server/producer"
	"github.com/peterP1998/notification-system/notification-server/server"
	"log"
)

func main() {
	err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}
	producer.InitProducer()
	fmt.Println("Hello, World")
	server.Init()
}
