package main

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-sender/server"
	"github.com/peterP1998/notification-system/notification-sender/consumer"
)

func main() {
	consumer.CreateSubscriber()
	fmt.Println("Hello, World")
	server.Init()
}