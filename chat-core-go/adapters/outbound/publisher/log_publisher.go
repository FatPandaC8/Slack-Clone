package publisher

import (
	"fmt"

	"chat-core-go/domain/message"
)

type LogPublisher struct{}

func (p *LogPublisher) Publish(msg message.Message) error {
	fmt.Printf("EVENT: message sent from %s: %s\n",msg.ConversationID(), msg.Content().Value()) // exactly like my broadcast server-client
	return nil
}