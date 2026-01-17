package application

import (
	"core/port/in"
	"core/port/out"
)

type JoinRoomUseCase struct {
	roomRepo       out.RoomRepository
	roomMemberRepo out.RoomMemberRepository
}

func NewJoinRoomUseCase(
	roomRepo out.RoomRepository,
	roomMemberRepo out.RoomMemberRepository,
) *JoinRoomUseCase {
	return &JoinRoomUseCase{
		roomRepo:       roomRepo,
		roomMemberRepo: roomMemberRepo,
	}
}

func (uc *JoinRoomUseCase) JoinRoom(cmd in.JoinRoomCommand) error {
	roomEntity, err := uc.roomRepo.FindByInviteCode(cmd.InviteCode)
	if err != nil {
		return err
	}

	return uc.roomMemberRepo.AddMember(
		roomEntity.ID(),
		cmd.UserID,
		"Member", // default of joining is to become a member :)
	)
}