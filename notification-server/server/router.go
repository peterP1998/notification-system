package server

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/notification-system/notification-server/controller"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	notificate := new(controller.NotificationController)

	router.POST("/notificate", notificate.SendNotification)
	return router

}