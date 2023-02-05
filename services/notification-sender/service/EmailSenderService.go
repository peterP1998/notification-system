package service

import (
	//"github.com/nikoksr/notify/service/mail"
	"github.com/peterP1998/notification-system/libs/notification/model"
)

type EmailSenderService struct {
}

func (EmailSenderService) SendNotification(notification *model.Notification) error {
	return nil
}
