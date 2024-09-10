package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func GenerateRandomToken(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}

func SendEmail(toEmail, subject, message string) error {
	from := mail.NewEmail("Your App", "no-reply@yourapp.com")
	to := mail.NewEmail("", toEmail)
	msg := mail.NewSingleEmail(from, subject, to, message, "")
	client := sendgrid.NewSendClient("YOUR_SENDGRID_API_KEY")
	_, err := client.Send(msg)
	return err
}
