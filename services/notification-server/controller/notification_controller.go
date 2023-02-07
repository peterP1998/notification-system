package controller

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/peterP1998/notification-system/notification-server/service"
)

type NotificationControllerInterface interface {
	SendNotification(ctx *gin.Context)
}

type NotificationController struct {
	notificationService service.NotificationServiceInterface
}

func NotificationControllerCreate(notificationService service.NotificationServiceInterface) NotificationControllerInterface {
	return &NotificationController{
		notificationService: notificationService,
	}
}

func (notificationController *NotificationController) SendNotification(ctx *gin.Context) {
	var notification model.Notification
	err := ctx.ShouldBindJSON(&notification)
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	err = notificationController.notificationService.PublishNotification(&notification)

	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(400, err)
		return
	}
}
