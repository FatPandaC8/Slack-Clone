package dto

import "chat-core-go/domain/user"

// DTOs are for changing state, write; queries on the other hand does not need DTO
type CreateConversationCommand struct {
	ConversationID 		string
	Members  			[]user.ID
}