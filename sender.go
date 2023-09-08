package emailsender

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	login    string
	password string
	host     string
	port     int
}

// NewEmailSender - create new email sender
func NewEmailSender(login, password, host string, port int) *EmailSender {
	return &EmailSender{
		login:    login,
		password: password,
		host:     host,
		port:     port,
	}
}

type Email struct {
	Theme string
	Msg   string
	To    []string
}

// Send - send email
func (es *EmailSender) Send(e Email) (err error) {

	m := gomail.NewMessage()

	m.SetHeader("From", es.login)
	m.SetHeader("To", e.To...)
	m.SetHeader("Subject", e.Theme)
	m.SetBody("text/plain", e.Msg)

	d := gomail.NewDialer(es.host, es.port, es.login, es.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
