package in

import "chat-core-go/application/dto"

type CreateConversationPort interface {
	Execute(cmd dto.CreateConversationCommand) (*dto.CreateConversationResult, error)
}