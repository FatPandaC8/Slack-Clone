package message

import (
	"chat-core-go/domain/valueobject"
	"time"
)

// Message is an ENTITY
type Message struct {
	id             valueobject.MessageID
	sender         valueobject.UserID
	conversationID valueobject.ConversationID
	content        valueobject.MessageContent
	createdAt      time.Time
}

// NewMessage creates a new message (factory method)
func NewMessage(
	id valueobject.MessageID,
	sender valueobject.UserID,
	conversationID valueobject.ConversationID,
	content valueobject.MessageContent,
) *Message {
	return &Message{
		id:             id,
		sender:         sender,
		conversationID: conversationID,
		content:        content,
		createdAt:      time.Now(),
	}
}

// ReconstructMessage recreates a message from persistence
func ReconstructMessage(
	id valueobject.MessageID,
	sender valueobject.UserID,
	conversationID valueobject.ConversationID,
	content valueobject.MessageContent,
	createdAt time.Time,
) *Message {
	return &Message{
		id:             id,
		sender:         sender,
		conversationID: conversationID,
		content:        content,
		createdAt:      createdAt,
	}
}

// Getters
func (m *Message) ID() valueobject.MessageID {
	return m.id
}

func (m *Message) Sender() valueobject.UserID {
	return m.sender
}

func (m *Message) ConversationID() valueobject.ConversationID {
	return m.conversationID
}

func (m *Message) Content() valueobject.MessageContent {
	return m.content
}

func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}

// Equals checks if two messages are the same
func (m *Message) Equals(other *Message) bool {
	if other == nil {
		return false
	}
	return m.id.Equals(other.id)
}