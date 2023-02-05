package service

import (
	"github.com/nikoksr/notify/service/mail"
	"github.com/peterP1998/libs/notification/model"
)

type EmailSenderService struct {
}

func (EmailSenderService) SendNotification(notification *model.Notification) {
	topic := buildTopic(notification.Type)
	producer.ProduceMessage(notification, topic)
}

func buildTopic(notificationType model.NotificationType) (string) {
	topicSuffix := "-notification-topic"
    topic := strings.ToLower(string(notificationType)) + topicSuffix
    
    return topic
}