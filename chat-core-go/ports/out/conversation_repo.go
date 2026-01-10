package out

import (
	"chat-core-go/domain/conversation"
)

type ConversationRepository interface {
	Load(id string) (*conversation.Conversation, error)
	Save(conv *conversation.Conversation) error
	FindByMember(userId string) ([]conversation.Conversation, error) // use a copy of conversation because, otherwise people can change the content of the conversation
	FindByInviteCode(code string) (*conversation.Conversation, error)
	GenerateID() string
}