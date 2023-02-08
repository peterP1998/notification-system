package service

import (
	"encoding/json"
	"fmt"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/peterP1998/notification-system/notification-sender/config"
	"github.com/peterP1998/notification-system/notification-sender/constants"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSSenderService struct {
}

func (SMSSenderService) SendNotification(notification *model.Notification) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.GetConfigProperty(constants.SMS_ACCOUNTID),
		Password: config.GetConfigProperty(constants.SMS_TOKEN),
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(notification.Receiver)
	params.SetFrom(config.GetConfigProperty(constants.SMS_FROM))
	params.SetBody(notification.Message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}
