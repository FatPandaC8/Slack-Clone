package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/conversation"
	"chat-core-go/ports/out"
)

type CreateConversation struct {
	conversations out.ConversationRepository
}

func NewCreateConversation(
	conversations out.ConversationRepository,
) *CreateConversation {
	return &CreateConversation{conversations: conversations}
}

func (uc *CreateConversation) Execute(cmd dto.CreateConversationCommand) error {
	conv, err := conversation.NewConversation(
		conversation.ID(cmd.ConversationID),
		cmd.Members,
	)
	if err != nil {
		return err
	}

	return uc.conversations.Save(conv) // TODO: sanity check
}