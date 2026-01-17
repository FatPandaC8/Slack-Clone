package in

import "chat-core-go/application/dto"

type SendMessagePort interface {
	Execute(dto.SendMessageCommand) error
}