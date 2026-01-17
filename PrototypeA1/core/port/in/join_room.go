package in

type JoinRoomCommand struct {
	UserID     string
	InviteCode string
}

type JoinRoomUseCase interface {
	JoinRoom(cmd JoinRoomCommand) error
}
