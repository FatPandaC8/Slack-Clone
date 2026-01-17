package in

type SendMessageCommand struct {
	UserID  string
	RoomID  string
	Content string
}

type SendMessageUseCase interface {
	SendMessage(cmd SendMessageCommand) error
}
