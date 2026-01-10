package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/conversation"
	"chat-core-go/ports/out"
	"math/rand"
)

type CreateConversation struct {
	conversations out.ConversationRepository
}

func NewCreateConversation(
	conversations out.ConversationRepository,
) *CreateConversation {
	return &CreateConversation{conversations: conversations}
}

func (uc *CreateConversation) Execute(cmd dto.CreateConversationCommand) (*dto.CreateConversationDTO, error) {
	id := uc.conversations.GenerateID()
	invite := generateInviteCode()

	conv, err := conversation.NewConversation(
		id, 
		cmd.Name,
		invite,
		cmd.CreatorID,
	)
	if err != nil {
		return nil, err
	}

	save_err := uc.conversations.Save(conv)
	if save_err != nil {
		return nil, save_err
	}

	return &dto.CreateConversationDTO{
		ID: id,
		InviteCode: invite,
		Name: cmd.Name,
	}, nil
}

func generateInviteCode() string { // random 6 char invite code (later generate by database index through base62 or base64)
    letters := "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
    b := make([]byte, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}