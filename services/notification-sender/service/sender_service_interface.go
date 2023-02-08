package service

import (
	"github.com/peterP1998/notification-system/libs/notification/model"
)

// This sender service interface is used to represent the different types
// of services, which are responsible for sending model.Notification.
type SenderServiceInterface interface {

	// The SendNotification method has to send notification
	// directly to the user, using third party library.
	SendNotification(notification *model.Notification) error
}
