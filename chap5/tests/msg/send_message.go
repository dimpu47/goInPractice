package msg

type Message struct {
	// ...
}

func (m *Message) Send(email, subject string, body []byte) error {
	// ...
	return nil
}

// Messages interface describes methods to use on Message struct type.
type Messager interface {
	Send(email, subject string, body []byte) error
}

// Passes that interface(m Messager) instead of message type
func Alert(m Messager, problem []byte) error {
	return m.Send("clark@example.com", "Critical Error", problem)
}
