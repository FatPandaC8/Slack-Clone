package room

import (
	"errors"
	"time"
)

var ErrAlreadyMember = errors.New("user already in room")
var ErrNotAdmin = errors.New("only admin can perform this action")

type Room struct {
	roomID         string
	name       string
	inviteCode string
	members    map[string]Member
	createdAt  time.Time
}

func NewRoom(id, name, inviteCode string, createdAt time.Time) *Room {
	return &Room{
		roomID:         id,
		name:       name,
		inviteCode: inviteCode,
		members:    make(map[string]Member),
		createdAt:  createdAt,
	}
}

func (r *Room) ID() string {
	return r.roomID
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) InviteCode() string {
	return r.inviteCode
}

func (r *Room) AddMember(userID string, role Role, joinedAt time.Time) error {
	if _, exists := r.members[userID]; exists {
		return ErrAlreadyMember
	}

	r.members[userID] = NewMember(userID, role, joinedAt)
	return nil
}

func (r *Room) IsMember(userID string) bool {
	_, ok := r.members[userID]
	return ok
}

func (r *Room) IsAdmin(userID string) bool {
	m, ok := r.members[userID]
	return ok && m.IsAdmin()
}
