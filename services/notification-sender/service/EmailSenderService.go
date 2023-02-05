package service

import (
	//"fmt"
	"net/smtp"
	"log"
	"github.com/peterP1998/notification-system/libs/notification/model"
)

type EmailSenderService struct {
}

func (EmailSenderService) SendNotification(notification *model.Notification) error {
	from := "petar.petrov220998@gmail.com"

    user := "petar.petrov220998@gmail.com"
    password := ""

    to := []string{
        "petar.petrov220998@gmail.com",
    }

    addr := "smtp.gmail.com:587"
    host := "smtp.gmail.com"

	// TODO create template for all notifications
    msg := []byte(
        "Subject: Notification\r\n\r\n" +
        "Samo levski Yisagi!\r\n")

    auth := smtp.PlainAuth("", user, password, host)

    err := smtp.SendMail(addr, auth, from, to, msg)

    if err != nil {
        log.Fatal(err)
    }

   return nil
}
