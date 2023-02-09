package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/peterP1998/notification-system/notification-server/service"
	"log"
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
		log.Printf("Error: %s", err)
		ctx.JSON(400, gin.H{"error": "Parameters are not ok.", "status": "Failed"})
		return
	}
	err = notificationController.notificationService.PublishNotification(&notification)

	if err != nil {
		log.Printf("Error: %s", err)
		ctx.JSON(400, gin.H{"error": err.Error(), "status": "Failed"})
		return
	}

	ctx.JSON(202, gin.H{"status": "Success"})
}
