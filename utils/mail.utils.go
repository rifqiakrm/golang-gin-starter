package utils

import (
	"gopkg.in/gomail.v2"

	"gin-starter/config"
)

func SendMail(cfg config.Config, to string, cc string, subject, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", cfg.SMTP.From)
	mailer.SetHeader("To", to)

	if cc != "" {
		mailer.SetAddressHeader("Cc", cc, "STARTER")
	}
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		cfg.SMTP.Host,
		cfg.SMTP.Port,
		cfg.SMTP.User,
		cfg.SMTP.Pass,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
