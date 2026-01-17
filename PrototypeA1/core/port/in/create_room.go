package in

type CreateRoomCommand struct {
	UserID string
	Name   string
}

type CreateRoomResult struct {
	RoomID     string
	InviteCode string
}

type CreateRoomUseCase interface {
	CreateRoom(cmd CreateRoomCommand) (*CreateRoomResult, error)
}
