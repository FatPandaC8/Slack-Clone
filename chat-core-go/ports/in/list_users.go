package in

import "chat-core-go/application/dto"

type ListUserPort interface {
	Execute(conversationId string) ([]*dto.UserDTO, error)
}