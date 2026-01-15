package out

import (
	"chat-core-go/domain/conversation"
	"chat-core-go/domain/valueobject"
)

type ConversationRepository interface {
	// Save persists a conversation (insert or update)
	Save(conv *conversation.Conversation) error
	
	// Load retrieves a conversation by ID
	Load(id valueobject.ConversationID) (*conversation.Conversation, error)
	
	// LoadByInviteCode retrieves a conversation by invite code
	LoadByInviteCode(code valueobject.InviteCode) (*conversation.Conversation, error)
	
	// FindByMember returns all conversations a user is a member of
	FindByMember(userID valueobject.UserID) ([]*conversation.Conversation, error)
	
	// GenerateID creates a new unique conversation ID
	GenerateID() valueobject.ConversationID
	
	// Exists checks if a conversation exists
	Exists(id valueobject.ConversationID) (bool, error)
}