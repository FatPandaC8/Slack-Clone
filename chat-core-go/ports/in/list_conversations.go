package in

import (
	"chat-core-go/domain/conversation"
	"chat-core-go/domain/user"
)

type ListConversationPort interface {
	Execute(userId user.ID) ([]conversation.Conversation, error)
}