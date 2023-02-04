package server

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/notification-system/notification-server/controller"
	"github.com/peterP1998/notification-system/notification-server/service"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	notificationService := service.NotificationService{}
	notificate := controller.NotificationControllerCreate(notificationService)

	router.POST("/notificate", notificate.SendNotification)
	return router

}