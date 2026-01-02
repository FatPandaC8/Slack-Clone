package dto

import "chat-core-go/domain/user"

type SendMessageCommand struct {
	MessageID 		string
	ConversationID 	string
	SenderID 		user.ID
	Text 			string
}