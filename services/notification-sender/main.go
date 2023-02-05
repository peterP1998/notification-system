package main

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-sender/consumer"
	"github.com/peterP1998/notification-system/notification-sender/server"
)

func main() {
	consumer.CreateSubscriber()
	fmt.Println("Hello, World")
	server.Init()
}
