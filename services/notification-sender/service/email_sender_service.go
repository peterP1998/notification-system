package service

import (
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/peterP1998/notification-system/notification-sender/config"
	"github.com/peterP1998/notification-system/notification-sender/constants"
	"net/smtp"
)

type EmailSenderService struct {
}

func (EmailSenderService) SendNotification(notification *model.Notification) error {

	to := []string{
		notification.Receiver,
	}

	msg := []byte(
		"Subject: Notification\r\n\r\n" +
			notification.Message + "\r\n")

	auth := smtp.PlainAuth("", config.GetConfigProperty(constants.EMAIL_FROM),
		config.GetConfigProperty(constants.EMAIL_PASSWORD),
		config.GetConfigProperty(constants.EMAIL_HOST))

	err := smtp.SendMail(config.GetConfigProperty(constants.EMAIL_ADDR),
		auth, config.GetConfigProperty(constants.EMAIL_FROM), to, msg)

	if err != nil {
		return err
	}

	return nil
}
