package in

import "chat-core-go/application/dto"

type RegisterUserPort interface {
	Execute(cmd dto.RegisterUserCommand) (*dto.RegisterUserResult, error)
}