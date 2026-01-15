// application/usecase/get_conversation.go
package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/ports/out"
	"errors"
)

// GetConversation is a query use case for retrieving conversation details
type GetConversation struct {
	conversations out.ConversationRepository
	messages      out.MessageRepository
	users         out.UserRepository
}

func NewGetConversation(
	conversations out.ConversationRepository,
	messages out.MessageRepository,
	users out.UserRepository,
) *GetConversation {
	return &GetConversation{
		conversations: conversations,
		messages:      messages,
		users:         users,
	}
}

// Execute performs the get conversation query
func (uc *GetConversation) Execute(query dto.GetConversationQuery) (*dto.GetConversationResult, error) {
	// 1. Validate principal
	if query.Principal == nil {
		return nil, errors.New("principal is required")
	}
	if query.Principal.IsExpired() {
		return nil, errors.New("principal expired")
	}
	
	// 2. Load conversation
	conv, err := uc.conversations.Load(query.ConversationID)
	if err != nil {
		return nil, errors.New("conversation not found")
	}
	
	// 3. Authorization: Check membership
	if !conv.HasMember(query.Principal.UserID()) {
		return nil, errors.New("not authorized: user is not a member")
	}
	
	// 4. Load messages
	messages, err := uc.messages.LoadByConversation(query.ConversationID)
	if err != nil {
		return nil, err
	}
	
	messageDTOs := make([]dto.MessageDTO, 0, len(messages))
	for _, msg := range messages {
		// Load sender info
		sender, err := uc.users.Load(msg.Sender())
		if err != nil {
			continue // Skip messages with invalid senders
		}
		
		messageDTOs = append(messageDTOs, dto.MessageDTO{
			MessageID:  msg.ID(),
			SenderID:   msg.Sender(),
			SenderName: sender.Name(),
			Content:    msg.Content(),
			CreatedAt:  msg.CreatedAt(),
		})
	}
	
	// 5. Load member info
	memberDTOs := make([]dto.MemberDTO, 0, conv.MemberCount())
	for _, memberID := range conv.Members() {
		member, err := uc.users.Load(memberID)
		if err != nil {
			continue // Skip invalid members
		}
		
		memberDTOs = append(memberDTOs, dto.MemberDTO{
			UserID: memberID,
			Name:   member.Name(),
		})
	}
	
	// 6. Return result
	return &dto.GetConversationResult{
		ConversationID: conv.ID(),
		Name:           conv.Name(),
		Members:        memberDTOs,
		Messages:       messageDTOs,
	}, nil
}