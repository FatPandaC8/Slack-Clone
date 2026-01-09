package out

import "chat-core-go/domain/message"

type MessageRepository interface {
	Save(msg message.Message) error
	Load(id string) (message.Message, error)
}
