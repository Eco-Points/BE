package utils

import (
	"eco_points/config"
	"log"

	"gopkg.in/gomail.v2"
)

type GomailUtilityInterface interface {
	SendEmail(points int, message, receiverEmail, receiverName string) error
}

type GomailUtility struct{}

func NewGomailUtility() GomailUtilityInterface {
	return &GomailUtility{}
}

func (gu *GomailUtility) SendEmail(points int, message, receiverEmail, receiverName string) error {
	emailKey := config.ImportGomailSetting()

	CONFIG_SMTP_HOST := emailKey.Host
	CONFIG_SMTP_PORT := emailKey.Port
	CONFIG_SENDER_NAME := emailKey.Name
	CONFIG_AUTH_EMAIL := emailKey.Email
	CONFIG_AUTH_PASSWORD := emailKey.Password

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", receiverEmail)
	mailer.SetHeader("Subject", "Eco Points : Perihal Pengajuan Deposit Sampah Anda")
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")

	return nil
}
