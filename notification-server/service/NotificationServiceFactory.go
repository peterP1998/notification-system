package service

import ( 
	"fmt"
	"github.com/peterP1998/notification-system/notification-server/model"
)

func GetNotificationService(notificationType model.NotificationType) (NotificationService, error) {
    if notificationType == "EMAIL" {
        return EmailNotificationService{}, nil
    }
    if notificationType == "SLACK" {
        return SlackNotificationService{}, nil
    }
    return nil, fmt.Errorf("Wrong type passed")
}