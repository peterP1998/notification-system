package service

import (
	"github.com/peterP1998/notification-system/libs/notification/model"
)

type NotificationServiceInterface interface {
	PublishNotification(notification *model.Notification) (error)
}