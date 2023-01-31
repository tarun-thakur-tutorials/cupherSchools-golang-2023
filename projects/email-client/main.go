package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"

	gomail "gopkg.in/mail.v2"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/cast"
)

func GenRandom() string {
	randomeCrypto, _ := rand.Prime(rand.Reader, 128)
	return cast.ToString(randomeCrypto)
}

func SendmailBySendGrid() {
	// sendgridKey := os.Getenv("SendgridKey")
	// sendgridKey := "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"
	sendgridKey := "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"
	from := mail.NewEmail(os.Getenv("JeremyName"), os.Getenv("JeremyEmail"))
	to := mail.NewEmail("tarun", "thakur.cs.tarun@gmail.com")
	subject := "Sending with Twilio SendGrid is Fun"
	plainTextContent := "and easy to do anywhere, even with Go"
	otp := GenRandom()
	htmlContent := fmt.Sprintf("<strong>OPT: %v</strong>", otp)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("plz enter the otp")

	var usrInput string
	fmt.Scanln(&usrInput)

	if usrInput == otp {
		fmt.Println("passed")
	} else {
		fmt.Println("failed")
	}

}

func SendEmailByGolangLib() {
	// Sender data.
	from := os.Getenv("MyEmail")
	password := os.Getenv("MyEmailPass")

	fmt.Println("from: ", from, " password: ", password)

	// Receiver email address.
	to := []string{
		"aayush.deep.contact@gmail.com",
		"22906.akanksha@gmail.com",
		"mnirmalkar745@gmail.com",
		"shivpaly2@gmai.com",
		"mshivam019@gmail.com",
		"nujgupta7906@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
func Gopkg() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "thakur.cs.tarun@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "aayush.deep.contact@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("MyEmail"), os.Getenv("MyEmailPass"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
func main() {
	SendmailBySendGrid()
	// Gopkg()
	// SendEmailByGolangLib()
}
