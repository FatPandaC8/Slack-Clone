package dto

import (
	"time"
)

type SendMessageCommand struct {
	MessageID 		string
	ConversationID 	string
	SenderID 		string
	Text 			string
	CreatedAt 		time.Time
}