package mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Mail struct {
	To      string
	Subject string
	Body    string
}

func (mail *Mail) BuildMessage() []byte {
	message := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	message += fmt.Sprintf("From: Health Centre Automation System IITK<%s>\r\n", sender)
	message += fmt.Sprintf("Subject: %s | Health Centre Automation System\r\n", mail.Subject)


	message += fmt.Sprintf("To: %s\r\n\r\n", mail.To)

	message += strings.Replace(mail.Body, "\n", "<br>", -1)
	message += "<br><br>--<br>Health Centre Automation System<br>"
	message += "Indian Institute of Technology Kanpur<br><br>"
	message += "<small>This is an auto-generated email. Please do not reply.</small>"

	return []byte(message)
}

func Service(mailQueue chan Mail) {
	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", sender, pass, host)

	for mail := range mailQueue {
		message := mail.BuildMessage()
		to := mail.To
		toSendArray := make([]string, 0)
		toSendArray = append(toSendArray, to)
		if err := smtp.SendMail(addr, auth, sender, toSendArray, message); err != nil {
			logrus.Errorf("Error sending mail: %v", to)
			logrus.Errorf("Error: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}
