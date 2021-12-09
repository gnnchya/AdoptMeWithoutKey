package email

import (
	"crypto/tls"
	"fmt"
	"github.com/gnnchya/AdoptMeWithoutKey/post/config"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	gomail "gopkg.in/mail.v2"
)

func SendAdoptEmail(email string, adopter domain.UserStruct) (err error) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Get().Email)

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Your Pet has been adopted!!!")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is your adopter\n"+"Name:"+adopter.Name+"\nPlease contact them by this Email: "+adopter.Email)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, config.Get().Email, config.Get().EmailPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SendFoundEmail(email string, founder domain.UserStruct) (err error) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Get().Email)

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Your Pet has been found!!!")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This person found your pet\n"+"Name:"+founder.Name+"\nPlease contact them by this Email: "+founder.Email)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, config.Get().Email, config.Get().EmailPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
