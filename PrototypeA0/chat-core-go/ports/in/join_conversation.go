package in

import "chat-core-go/application/dto"

type JoinConversationPort interface {
	Execute(cmd dto.JoinConversationCommand) error
}