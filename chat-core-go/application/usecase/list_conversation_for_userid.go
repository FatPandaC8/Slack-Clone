package usecase

import (
	"chat-core-go/domain/conversation"
	"chat-core-go/ports/out"
)

type ListConversations struct {
	repo out.ConversationRepository
}

func NewListConversations(r out.ConversationRepository) *ListConversations {
	return &ListConversations{
		repo: r,
	}
}

func (uc *ListConversations) Execute(uid string) ([]conversation.Conversation, error) {
	return uc.repo.FindByMember(uid)
}