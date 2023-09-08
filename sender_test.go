package emailsender_test

import (
	emailsender "email-sender"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateEmailSender() *emailsender.EmailSender {
	return emailsender.NewEmailSender("login", "password", "host", 425)
}

func TestSend(t *testing.T) {
	es := CreateEmailSender()
	err := es.Send(emailsender.Email{
		Theme: "Test message",
		Msg:   "This is a test message",
		To:    []string{"test_addr@test.com"},
	})
	assert.NoError(t, err)
}
