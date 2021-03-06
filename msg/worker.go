package msg

import (
	"context"
	"log"
)

// Request stores SMS details that will be used in NewMessage.
type Request struct {
	Sender    string
	Recipient string
	Text      string
}

// Messenger is used to send text messages.
type Messenger interface {
	SendText(sender, msisdn, text string) error
}

// StartSendingMessages starts background worker that sends short messages.
func StartSendingMessages(ctx context.Context, mc chan Request, m Messenger) {
	var req Request

	for {
		select {
		case <-ctx.Done():
			return
		case req = <-mc:
			err := m.SendText(req.Sender, req.Recipient, req.Text)
			if err != nil {
				log.Println("Failed to send SMS, error:", err)
			}
		}
	}
}
