package main

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"io/ioutil"
)

/*
	This function will not work as google has blocked less secure apps access
	unless the sender config to allow such operation
	Check this link for more details: https://support.google.com/accounts/answer/6010255
*/
func main() {
	m := gomail.NewMessage()

	// Set email sender
	m.SetHeader("From", "sender@example.com")

	// Set email receivers
	m.SetHeader("To", "personA@example.com", "personB@example.com")

	// Set email subject
	subject := "Testing"
	m.SetHeader("Subject", subject)

	// Set email body. Http email is supported but here we use plaintext obtained from a txt
	body, _ := ioutil.ReadFile("template.txt")
	m.SetBody("text/plain", string(body))

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "sender@example.com", "<Email Password>")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
