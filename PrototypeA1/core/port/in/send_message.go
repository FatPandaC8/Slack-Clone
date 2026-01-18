package in

import valueobject "core/domain/valueobject/user"

type SendMessageCommand struct {
	UserID  valueobject.UserID
	RoomID  string
	Content string
}

type SendMessageUseCase interface {
	SendMessage(cmd SendMessageCommand) error
}
