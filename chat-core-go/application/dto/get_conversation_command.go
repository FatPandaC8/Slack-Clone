package dto

import (
	"chat-core-go/domain/identity"
	"chat-core-go/domain/valueobject"
	"time"
)

// GetConversationQuery contains input for the query
type GetConversationQuery struct {
	Principal      *identity.Principal
	ConversationID valueobject.ConversationID
}

// GetConversationResult contains the conversation details
type GetConversationResult struct {
	ConversationID valueobject.ConversationID
	Name           valueobject.ConversationName
	Members        []MemberDTO
	Messages       []MessageDTO
}

// MemberDTO represents a conversation member
type MemberDTO struct {
	UserID valueobject.UserID
	Name   valueobject.UserName
}

// MessageDTO represents a message
type MessageDTO struct {
	MessageID  valueobject.MessageID
	SenderID   valueobject.UserID
	SenderName valueobject.UserName
	Content    valueobject.MessageContent
	CreatedAt  time.Time
}