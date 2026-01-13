package in

import "chat-core-go/application/dto"

type LoginUserPort interface {
	Execute(cmd dto.LoginUserCommand) (*dto.LoginResultDTO, error)
}