package server

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/notification-system/notification-server/controller"
	"github.com/peterP1998/notification-system/notification-server/producer"
	"github.com/peterP1998/notification-system/notification-server/service"
	"log"
)

func InitRouter(kafkaHost string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	producerClient := producer.CreateProducer(kafkaHost)
	log.Printf("%s", producerClient)
	notificationService := service.NotificationServiceCreate(producerClient)
	notification := controller.NotificationControllerCreate(notificationService)

	router.POST("/notification", notification.SendNotification)
	return router

}
