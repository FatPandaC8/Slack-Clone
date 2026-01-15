package dto

import (
	"chat-core-go/domain/identity"
	"chat-core-go/domain/valueobject"
)

type SendMessageCommand struct {
	Principal      *identity.Principal
	ConversationID valueobject.ConversationID
	Content        valueobject.MessageContent
}