package in

import "chat-core-go/application/dto"

type CreateUserPort interface {
	Execute(cmd dto.CreateUserCommand) (*dto.UserDTO, error)
}