package service

// The purpose of this interface is to send message using the correct service
type SenderServiceFacadeInterface interface {

	// This method takes as parameter byte array, which represents the message
	// we want to send.
	SendNotification(message []byte) error
}
