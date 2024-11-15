package email

type EmailClient interface {
	SendEmail(email string, subject string, body string) error
}

type EmailService interface {
	SendEmail(email string, subject string, body string) error
	SendMagicLink(email string, link string) error
}
