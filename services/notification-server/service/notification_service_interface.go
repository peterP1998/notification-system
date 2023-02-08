package service

type NotificationServiceInterface interface {
	PublishNotification(notification *model.Notification) (error)
}