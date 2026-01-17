package out

import "chat-core-go/domain/message"

type MessagePublisher interface {
	Publish(message.Message) error
}
