package main

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-server/server"
	"github.com/peterP1998/notification-system/notification-server/producer"
)

func main() {
	producer.InitProducer()
	fmt.Println("Hello, World")
	server.Init()
}