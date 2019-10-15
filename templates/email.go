package main

import (
	"bytes"
	"net/smtp"
	"strconv"
	"text/template"
)

type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}

func main() {
	message := &EmailMessage{
		From:    "sajioloye@gmail.com",
		To:      []string{"adams.banjo@max.ng", "folabibanjo@gmail.com"},
		Subject: "Emails with Golang",
		Body:    "How far can our imaginations carry us?",
	}

	//* populate a buffer with the rendered message text from the template
	var body bytes.Buffer
	t.Execute(&body, message)

	//* set up the smtp mail client
	authCreds := &EmailCredentials{
		Username: "",
		Password: "",
		Server:   "smtp.mailtrap.io",
		Port:     2525,
	}

	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
	)

	//* sends the email
	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		//* the bytes from the message buffer are passed in when the message is sent
		body.Bytes())
}
