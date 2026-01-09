package dto

import (
	"chat-core-go/domain/user"
	"time"
)

type SendMessageCommand struct {
	MessageID 		string
	ConversationID 	string
	SenderID 		user.ID
	Text 			string
	CreatedAt 		time.Time
}