package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/notification-system/notification-server/service"
	"github.com/peterP1998/notification-system/notification-server/model"
)

type NotificationControllerInterface interface {
	SendNotification(ctx *gin.Context)
}

type NotificationController struct {

}

func (artistController *NotificationController) SendNotification(ctx *gin.Context) {
	var notification model.Notification
	err := ctx.ShouldBindJSON(&notification)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	notificationService, _ := service.GetNotificationService(notification.Type)
	notificationService.PublishNotification(&model.Notification{})
}