package email

import "github.com/resend/resend-go/v2"

type resendEmailClient struct {
	c *resend.Client
}

func NewResendEmailClient(apiKey string) EmailClient {
	client := resend.NewClient(apiKey)
	return &resendEmailClient{c: client}
}

func (e *resendEmailClient) SendEmail(email string, subject string, body string) error {
	params := &resend.SendEmailRequest{
		From:    "",
		To:      []string{email},
		Subject: subject,
		Html:    body,
		Text:    "",
	}
	_, err := e.c.Emails.Send(params)
	return err
}
