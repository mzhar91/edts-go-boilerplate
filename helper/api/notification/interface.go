package notification

type ServiceNotification interface {
	SendEmail(param *EmailRequest, token string) (*EmailResponse, error)
}
