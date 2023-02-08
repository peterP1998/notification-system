package service

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"log"
)

type SlackSenderService struct {
}

func (SlackSenderService) SendNotification(notification *model.Notification) error {
	log.Printf("Sending slack notification %v", notification)
	webhookUrl := notification.Receiver

	payload := slack.Payload{
		Text: notification.Message,
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		log.Printf("Erorr is thrown in the slack notification service: %v", err)
		return err[0]
	}

	log.Printf("Slack notification send notification successfully %v", notification)
	return nil
}
