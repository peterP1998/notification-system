package service

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/peterP1998/notification-system/libs/notification/model"
)

type SlackSenderService struct {
}

func (SlackSenderService) SendNotification(notification *model.Notification) error {
	webhookUrl := notification.Receiver

	payload := slack.Payload{
		Text: notification.Message,
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	return nil
}
