package service

import (
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/peterP1998/notification-system/notification-server/producer"
	"strings"
)

const topicSuffix = "-notification-topic"

type NotificationServiceInterface interface {
	PublishNotification(notification *model.Notification)
}

type NotificationService struct {
}

func (NotificationService) PublishNotification(notification *model.Notification) {
	topic := buildTopic(notification.Type)
	producer.ProduceMessage(notification, topic)
}

func buildTopic(notificationType model.NotificationType) string {
	topic := strings.ToLower(string(notificationType)) + topicSuffix

	return topic
}
