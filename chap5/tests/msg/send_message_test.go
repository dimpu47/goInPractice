package msg

import (
	"testing"
)

type MockMessage struct {
	email, subject string
	body           []byte
}

// MockMessage implements Messager
func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage)         // creates new MockMessage instance
	body := []byte("Critical Error") // mock property

	Alert(msgr, body) // calling Alert func with mock vals

	// Accessing MockMessage Properties for verfication.
	if msgr.subject != "Critical Error" {
		t.Errorf("Expected 'Critical Error', Got '%s'", msgr.subject)
	}
}
