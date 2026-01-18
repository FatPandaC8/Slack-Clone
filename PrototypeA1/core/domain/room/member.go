package room

import (
	valueobject "core/domain/valueobject/user"
	"time"
)

type Member struct {
	userID 		valueobject.UserID
	role 		Role
	joinedAt 	time.Time
}

func NewMember(userID valueobject.UserID, role Role, joinedAt time.Time) Member {
	return Member{
		userID:   userID,
		role:     role,
		joinedAt: joinedAt,
	}
}

func (m Member) UserID() string {
	return m.userID.String()
}

func (m Member) Role() Role {
	return m.role
}

func (m Member) IsAdmin() bool {
	return m.role == RoleAdmin
}