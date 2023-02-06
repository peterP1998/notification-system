package service

import (
	"encoding/json"
	"fmt"
	"github.com/peterP1998/notification-system/libs/notification/model"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSSenderService struct {
}

func (SMSSenderService) SendNotification(notification *model.Notification) error {

	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "f2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(notification.Receiver)
	params.SetFrom("+15017250604")
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
