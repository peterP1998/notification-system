package service

import (
	"encoding/json"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/peterP1998/notification-system/notification-server/producer"
	"strings"
	"log"
	"errors"
	//"regexp"
)

const (
	TOPCI_SUFFIX = "-notification-topic"
	REGEX_FOR_MESSAGE = `^[A-Za-z0-9_-]*$`
)

type NotificationService struct {
	producerClient producer.ProducerInterface
}

func NotificationServiceCreate(producerClient producer.ProducerInterface) NotificationServiceInterface {
	return &NotificationService{
		producerClient: producerClient,
	}
}

func (ns NotificationService) PublishNotification(notification *model.Notification) (error){
	topic := buildTopic(notification.Type)
	err := validateNotification(notification)
	if err != nil {
		return err
	}
	byteNotification, err := json.Marshal(notification)

	if err != nil {
		return err
	}

	log.Printf("Publishing notification")
	ns.producerClient.PublishNotification(byteNotification, topic)

	return nil
}

func buildTopic(notificationType model.NotificationType) string {
	topic := strings.ToLower(string(notificationType)) + TOPCI_SUFFIX
	return topic
}

func validateNotification(notification *model.Notification) (error) {
	if len(notification.Receiver) == 0 {
		return errors.New("receiver is empty")
	}

	if len(notification.Message) == 0 {
		return errors.New("message is empty")
	}

	if len(notification.Message) > 50 {
		return errors.New("message is longer than 50 chars")
	}

	/*if regexp.MustCompile(REGEX_FOR_MESSAGE).MatchString(notification.Message) {
		return errors.New("message is not following the correct format")
	}*/

	return nil
}


