package in

import "chat-core-go/application/dto"

type CreateChannelPort interface {
	Execute(cmd dto.CreateChannelCommand) error
}