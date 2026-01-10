package dto

import (
	"time"
)

// DTOs are for changing state, write; queries on the other hand does not need DTO
type CreateConversationCommand struct {
	Name 				string
	CreatorID 			string
	CreatedAt 			time.Time
}

type CreateConversationDTO struct {
	ID 					string
	InviteCode 			string
	Name 				string
}