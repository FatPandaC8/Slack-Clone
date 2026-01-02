package conversation

import (
	"chat-core-go/domain/message"
	"chat-core-go/domain/user"
)

type Conversation interface {
	ID() ID
	CanSend(user.ID) bool
	AddMessage(message.Message)
}