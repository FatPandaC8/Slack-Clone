package dto

import (
	"chat-core-go/domain/user"
	"time"
)

// DTOs are for changing state, write; queries on the other hand does not need DTO
type CreateConversationCommand struct {
	ConversationID 		string
	Members  			[]user.ID
	CreatedAt 			time.Time
}