package model

type NotificationType string

const (
	EMAIL NotificationType = "EMAIL"
	SLACK NotificationType = "SLACK"
	SMS   NotificationType = "SMS"
)

type Notification struct {
	Id       int
	Receiver string
	Message  string
	Type     NotificationType
}
