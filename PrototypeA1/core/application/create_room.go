package application

import (
	"core/domain/room"
	"core/port/in"
	"core/port/out"
	"time"

	"github.com/google/uuid"
)

type CreateRoomUseCase struct {
	roomRepo 	out.RoomRepository
	roomMemberRepo out.RoomMemberRepository
}

func NewCreateRoomUseCase(
	roomRepo out.RoomRepository,
	roomMemberRepo out.RoomMemberRepository,
) *CreateRoomUseCase {
	return &CreateRoomUseCase{
		roomRepo:       roomRepo,
		roomMemberRepo: roomMemberRepo,
	}
}

func (uc *CreateRoomUseCase) CreateRoom(cmd in.CreateRoomCommand) (*in.CreateRoomResult, error) {
	roomID := uuid.NewString()
	inviteCode := uuid.NewString()[:8]

	roomEntity := room.NewRoom(
		roomID,
		cmd.Name,
		inviteCode,
		time.Now(),
	)

	if err := uc.roomRepo.Save(roomEntity); err != nil {
		return nil, err
	}

	if err := uc.roomMemberRepo.AddMember(
		roomID,
		cmd.UserID,
		string(room.RoleAdmin),
	); err != nil {
		return nil, err
	}

	return &in.CreateRoomResult{
		RoomID: roomID,
		InviteCode: inviteCode,
	}, nil
}