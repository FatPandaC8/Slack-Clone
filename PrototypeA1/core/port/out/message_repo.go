package out

import "core/domain/message"

type MessageRepository interface {
	Save(msg *message.Message) error
	FindByRoom(roomID string) ([]*message.Message, error)
}