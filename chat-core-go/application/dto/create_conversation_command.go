package dto

import (
	"chat-core-go/domain/identity"
	"chat-core-go/domain/valueobject"
)

// DTOs are for changing state, write; queries on the other hand does not need DTO
type CreateConversationCommand struct {
	Principal *identity.Principal
	Name      valueobject.ConversationName
}

type CreateConversationResult struct {
	ConversationID valueobject.ConversationID
	Name           valueobject.ConversationName
	InviteCode     valueobject.InviteCode
}