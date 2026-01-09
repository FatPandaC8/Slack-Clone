package in

import (
	"chat-core-go/domain/conversation"
)

type ListConversationPort interface {
	Execute(userId string) ([]conversation.Conversation, error)
}