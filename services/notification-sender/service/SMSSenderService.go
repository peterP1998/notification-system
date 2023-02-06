package service

import (
	"github.com/peterP1998/notification-system/libs/notification/model"
)

type SMSSenderService struct {
}

func (SMSSenderService) SendNotification(notification *model.Notification) error {

	return nil
}
