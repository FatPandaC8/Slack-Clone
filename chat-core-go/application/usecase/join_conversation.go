package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/ports/out"
)

type JoinConversation struct {
	convs out.ConversationRepository
}

func NewJoinConversation(repo out.ConversationRepository) *JoinConversation {
	return &JoinConversation{
		convs: repo,
	}
}

func (uc *JoinConversation) Execute(cmd dto.JoinConversationCommand) error {
	conv, err := uc.convs.FindByInviteCode(cmd.InviteCode)
	if err != nil {
		return err
	}

	conv.AddMember(cmd.UserID)
	return uc.convs.Save(conv)
}