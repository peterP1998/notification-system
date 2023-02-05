package service

import (
	"github.com/peterP1998/notification-system/libs/notification/model")

type SenderServiceInterface interface {
	SendNotification(notification *model.Notification)
}