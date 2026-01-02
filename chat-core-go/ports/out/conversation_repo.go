package out

import "chat-core-go/domain/conversation"

type ConversationRepository interface {
	Load(id string) (*conversation.Conversation, error)
	Save(conv *conversation.Conversation) error
}