package usecase

import (
	"chat-core-go/domain/conversation"
	"chat-core-go/ports/out"
)

type GetConversation struct {
	conversations out.ConversationRepository
}

func NewGetConversation(repo out.ConversationRepository) *GetConversation {
	return &GetConversation{
		conversations: repo,
	}
}

func (uc *GetConversation) Execute(id string) (*conversation.Conversation, error) {
	return uc.conversations.Load(id)
}