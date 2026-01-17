package out

import "core/domain/room"

type RoomRepository interface {
	Save(room *room.Room) error
	FindByID(roomID string) (*room.Room, error)
	FindByInviteCode(inviteCode string) (*room.Room, error)
}