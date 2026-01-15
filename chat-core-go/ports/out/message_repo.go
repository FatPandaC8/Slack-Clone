package out

import (
	"chat-core-go/domain/message"
	"chat-core-go/domain/valueobject"
)

type MessageRepository interface {
	// Save persists a message
	Save(msg *message.Message) error
	
	// Load retrieves a message by ID
	Load(id valueobject.MessageID) (*message.Message, error)
	
	// LoadByConversation retrieves all messages in a conversation
	LoadByConversation(convID valueobject.ConversationID) ([]*message.Message, error)
	
	// GenerateID creates a new unique message ID
	GenerateID() valueobject.MessageID
	
	// Delete removes a message (soft or hard delete)
	Delete(id valueobject.MessageID) error
}