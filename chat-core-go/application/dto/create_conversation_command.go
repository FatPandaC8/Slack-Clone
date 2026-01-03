package dto

import "chat-core-go/domain/user"

type CreateConversationCommand struct {
	ConversationID 		string
	Members  			[]user.ID
}