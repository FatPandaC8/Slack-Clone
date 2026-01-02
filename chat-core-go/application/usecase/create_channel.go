package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/conversation"
	"chat-core-go/ports/out"
)

type CreateChannel struct {
	conversations out.ConversationRepository
}

func NewCreateChannel(
	conversations out.ConversationRepository,
) *CreateChannel {
	return &CreateChannel{conversations: conversations}
}

func (uc *CreateChannel) Execute(cmd dto.CreateChannelCommand) error {
	channel := conversation.NewChannel(
		conversation.ID(cmd.ChannelID),
		cmd.Members,
	)

	return uc.conversations.Save(channel)
}