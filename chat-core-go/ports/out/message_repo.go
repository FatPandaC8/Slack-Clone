package out

import "chat-core-go/domain/message"

type MessageRepository interface {
	Save(message.Message) error
}
