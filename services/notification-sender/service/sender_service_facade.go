package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/peterP1998/notification-system/libs/notification/model"
)

func SendNotification(message []byte) error {

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
		fmt.Println(err)
		return err
	}

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
		return nil, errors.New("can't find notification sender of this type")
	}

	return senderService, nil
}
