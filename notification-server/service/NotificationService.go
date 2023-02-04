package service

import (
	"github.com/peterP1998/notification-system/notification-server/model"
)

type NotificationService interface {
	PublishNotification(notification *model.Notification) (error)
}