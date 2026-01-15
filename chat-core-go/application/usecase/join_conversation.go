// application/usecase/join_conversation.go
package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/ports/out"
	"errors"
)

// JoinConversation is a use case for joining conversations via invite code
type JoinConversation struct {
	conversations out.ConversationRepository
	users         out.UserRepository
}

func NewJoinConversation(
	conversations out.ConversationRepository,
	users out.UserRepository,
) *JoinConversation {
	return &JoinConversation{
		conversations: conversations,
		users:         users,
	}
}

// Execute performs the join conversation use case
func (uc *JoinConversation) Execute(cmd dto.JoinConversationCommand) error {
	// 1. Validate principal
	if cmd.Principal == nil {
		return errors.New("principal is required")
	}
	if cmd.Principal.IsExpired() {
		return errors.New("principal expired")
	}
	
	// 2. Verify user exists
	userExists, err := uc.users.Exists(cmd.Principal.UserID())
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user not found")
	}
	
	// 3. Find conversation by invite code
	conv, err := uc.conversations.LoadByInviteCode(cmd.InviteCode)
	if err != nil {
		return errors.New("invalid invite code")
	}
	
	// 4. Add user as member (idempotent - won't fail if already member)
	err = conv.AddMember(cmd.Principal.UserID())
	if err != nil {
		return err
	}
	
	// 5. Persist changes
	err = uc.conversations.Save(conv)
	if err != nil {
		return err
	}
	
	return nil
}