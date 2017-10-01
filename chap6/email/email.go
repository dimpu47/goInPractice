package main

import (
	"bytes"
	"net/smtp"
	"strconv"
	"text/template"
)

type EmailMessage struct {
	From, Subject, Body string
	To	[]string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port	int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
{{.Body}}
`

var t *template.Template

func init() {
	t = template.New("Email")
	t.Parse(emailTemplate)
}

func main() {
	message := &EmailMessage{
		From: "me@example.com",
		Subject: "A test",
		Body: "Lean",
		To: []string{"gaurav.chdry47@gmail.com"},
	}

	var buf bytes.Buffer
	t.Execute(buf, message)

	authCreds := &EmailCredentials{
		Username: "",
		Password: "",
		Server: "smtp.example.com",
		Port: 25,
	}
	auth := smtp.PlainAuth("", 
		authCreds.Username,
		authCreds.Password, 
		authCreds.Server)
	smtp.SendMail(authCreds.Server + strconv.Itoa(authCreds.Port), 
					auth,
					message.From,
					message.To,
					buf.Bytes())
}
