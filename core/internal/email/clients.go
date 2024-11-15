package email

import (
	"fmt"
	"net/smtp"
)

type smptEmailClient struct {
	c    smtp.Auth
	host string
	port int
}

func NewSMPTEmailClient(user, password, host string, port int) EmailClient {
	auth := smtp.PlainAuth("", user, password, host)
	return &smptEmailClient{c: auth, host: host, port: port}
}

func (e *smptEmailClient) SendEmail(email string, subject string, body string) error {
	err := smtp.SendMail(fmt.Sprintf("%s:%d", e.host, e.port), e.c, email, []string{email}, []byte("Subject: "+subject+"\r\n\r\n"+body))
	return err
}
