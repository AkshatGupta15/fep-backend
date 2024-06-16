package mail

func GenerateMail(to, subject, body string) Mail {
	return Mail{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}
