package email

type emailService struct {
	c EmailClient
}

// SendEmail implements EmailService.
func (e *emailService) SendEmail(email string, subject string, body string) error {
	return e.c.SendEmail(email, subject, body)
}

// SendMagicLink implements EmailService.
func (e *emailService) SendMagicLink(email string, link string) error {
	return e.SendEmail(email, "Magic Link", link)
}

func NewEmailService(c EmailClient) EmailService {
	return &emailService{
		c: c,
	}
}
