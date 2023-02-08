package service

import (
	"encoding/json"
	"errors"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"log"
)

type SenderServiceFacade struct {
}

func (senderServiceFacade SenderServiceFacade) SendNotification(message []byte) error {

	log.Print("Sending notification")
	notification, err := convertToNotification(message)
	if err != nil {
		return err
	}

	senderService, err := senderServiceFactory(notification)

	if err != nil {
		return err
	}

	err = senderService.SendNotification(&notification)

	if err != nil {
		log.Print("Error is thrown in the ")
		return err
	}

	log.Print("Sending notification finished successfull")

	return nil
}

func convertToNotification(message []byte) (model.Notification, error) {
	var notification model.Notification
	err := json.Unmarshal(message, &notification)
	if err != nil {
		return model.Notification{}, err
	}

	return notification, nil
}

func senderServiceFactory(notification model.Notification) (SenderServiceInterface, error) {
	var senderService SenderServiceInterface
	if notification.Type == "EMAIL" {
		senderService = EmailSenderService{}
	} else if notification.Type == "SLACK" {
		senderService = SlackSenderService{}
	} else if notification.Type == "SMS" {
		senderService = SMSSenderService{}
	} else {
		return nil, errors.New("can't find notification sender service of this type")
	}

	return senderService, nil
}
