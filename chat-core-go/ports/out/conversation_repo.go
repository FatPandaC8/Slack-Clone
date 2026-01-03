package out

import (
	"chat-core-go/domain/conversation"
	"chat-core-go/domain/user"
)

type ConversationRepository interface {
	Load(id string) (*conversation.Conversation, error)
	Save(conv *conversation.Conversation) error
	FindByMember(userId user.ID) ([]conversation.Conversation, error) // use a copy of conversation because, otherwise people can change the content of the conversation
}