package service

import (
	"log"
	"testing"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/stretchr/testify/assert"
)

type ProducerClientMock struct {
}

func (p ProducerClientMock) PublishNotification(notification []byte, topic string) {
	log.Printf("Mock publish notification")
}

func TestPublishNotification(t *testing.T) {
	var ns NotificationService = NotificationService{ProducerClientMock{}}
	var notification = model.Notification{0, "test", "test31", "EMAIL"}
	
	err := ns.PublishNotification(&notification)
	assert.Equal(t, nil, err, "Validating is not working correctly")
	
	notification = model.Notification{0, "", "test31", "EMAIL"}
	
	err = ns.PublishNotification(&notification)
	assert.Equal(t, "receiver is empty", err.Error(), "Validating is not working correctly")

	notification = model.Notification{0, "test", "", "EMAIL"}
	
	err = ns.PublishNotification(&notification)
	assert.Equal(t, "message is empty" , err.Error(), "Validating is not working correctly")

	notification = model.Notification{0, "test", 
	       "testtesttesttesttestvtesttesttesttesttesttesttesttesttest", "EMAIL"}
	
	err = ns.PublishNotification(&notification)
	assert.Equal(t, "message is longer than 50 chars" , err.Error(), "Validating is not working correctly")
}

