package service

import (
	"fmt"
	"github.com/peterP1998/notification-system/notification-server/model"
)

type SlackNotificationService struct {
}

func (SlackNotificationService) PublishNotification(notification *model.Notification) (error) {
	fmt.Println("Levski2")
	return nil
}