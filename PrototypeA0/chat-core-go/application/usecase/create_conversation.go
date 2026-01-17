// application/usecase/create_conversation.go
package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/conversation"
	"chat-core-go/domain/valueobject"
	"chat-core-go/ports/out"
	"errors"
	"math/rand"
)

// CreateConversation is a use case for creating conversations
type CreateConversation struct {
	conversations out.ConversationRepository
}

func NewCreateConversation(
	conversations out.ConversationRepository,
) *CreateConversation {
	return &CreateConversation{
		conversations: conversations,
	}
}

// Execute performs the create conversation use case
func (uc *CreateConversation) Execute(cmd dto.CreateConversationCommand) (*dto.CreateConversationResult, error) {
	// 1. Validate principal
	if cmd.Principal == nil {
		return nil, errors.New("principal is required")
	}
	if cmd.Principal.IsExpired() {
		return nil, errors.New("principal expired")
	}
	
	// 2. Generate IDs
	conversationID := uc.conversations.GenerateID()
	inviteCode := generateInviteCode()
	
	// 3. Create conversation aggregate
	conv, err := conversation.NewConversation(
		conversationID,
		cmd.Name,
		inviteCode,
		cmd.Principal.UserID(),
	)
	if err != nil {
		return nil, err
	}
	
	// 4. Persist
	err = uc.conversations.Save(conv)
	if err != nil {
		return nil, err
	}
	
	// 5. Return result
	return &dto.CreateConversationResult{
		ConversationID: conversationID,
		Name:           cmd.Name,
		InviteCode:     inviteCode,
	}, nil
}

// generateInviteCode creates a random 6-character invite code
func generateInviteCode() valueobject.InviteCode {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return valueobject.MustInviteCode(string(b))
}