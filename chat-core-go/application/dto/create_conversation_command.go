package dto

import (
	"time"
)

// DTOs are for changing state, write; queries on the other hand does not need DTO
type CreateConversationCommand struct {
	ConversationID 		string
	Members  			[]string
	CreatedAt 			time.Time
}