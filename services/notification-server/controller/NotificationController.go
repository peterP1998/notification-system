package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/notification-system/notification-server/service"
	"github.com/peterP1998/libs/notification/model"
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
		fmt.Println(err)
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	notificationController.notificationService.PublishNotification(&notification)
}