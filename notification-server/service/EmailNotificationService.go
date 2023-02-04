package service

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-server/model"
)

type EmailNotificationService struct {
}

func (EmailNotificationService) PublishNotification(notification *model.Notification) (error) {
	fmt.Println("Levski")
	return nil
}