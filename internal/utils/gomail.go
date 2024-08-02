package utils

import (
	"eco_points/config"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

type GomailUtilityInterface interface {
	SendEmail(points int, receiverEmail, receiverName string) error
}

type GomailUtility struct{}

func NewGomailUtility() GomailUtilityInterface {
	return &GomailUtility{}
}

func (gu *GomailUtility) SendEmail(points int, receiverEmail, receiverName string) error {
	emailKey := config.ImportGomailSetting()

	CONFIG_SMTP_HOST := emailKey.Host
	CONFIG_SMTP_PORT := emailKey.Port
	CONFIG_SENDER_NAME := emailKey.Name
	CONFIG_AUTH_EMAIL := emailKey.Email
	CONFIG_AUTH_PASSWORD := emailKey.Password
	message := "Halo, <b>" + receiverName + "</b><br/> Deposit sampahmu sudah diverifikasi oleh admin keren dari tim Eco Points. <br/>Selamat kamu mendapatkan <b>" + strconv.Itoa(points) + " poin</b> "

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", receiverEmail)
	mailer.SetHeader("Subject", "Eco Points Regards")
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
