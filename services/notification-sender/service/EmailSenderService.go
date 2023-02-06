package service

import (
	"fmt"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"net/smtp"
)

type EmailSenderService struct {
}

// TODO refactor
func (EmailSenderService) SendNotification(notification *model.Notification) error {
	from := "petar.petrov220998@gmail.com"

	user := "petar.petrov220998@gmail.com"
	password := "" // TODO store password in secret file

	to := []string{
		notification.Receiver,
	}

	addr := "smtp.gmail.com:587"
	host := "smtp.gmail.com"

	// TODO create template for all notifications
	msg := []byte(
		"Subject: Notification\r\n\r\n" +
			notification.Message + "\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		return err
	}

	return nil
}
