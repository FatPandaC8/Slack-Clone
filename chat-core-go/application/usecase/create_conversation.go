package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/conversation"
	"chat-core-go/ports/out"
)

type CreateConversation struct {
	conversations out.ConversationRepository
}

func NewCreateChannel(
	conversations out.ConversationRepository,
) *CreateConversation {
	return &CreateConversation{conversations: conversations}
}

func (uc *CreateConversation) Execute(cmd dto.CreateChannelCommand) error {
	conv, err := conversation.NewConversation(
		conversation.ID(cmd.ChannelID),
		cmd.Members,
	)
	if err != nil {
		return err
	}

	return uc.conversations.Save(conv) // TODO: sanity check
}