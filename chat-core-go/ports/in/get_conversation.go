package in

import (
	"chat-core-go/application/dto"
)

type GetConversationPort interface {
	Execute(conversationID string) (*dto.ConversationDTO, error)
}