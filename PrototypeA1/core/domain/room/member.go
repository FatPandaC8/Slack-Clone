package room

import "time"

type Member struct {
	userID 		string
	role 		Role
	joinedAt 	time.Time
}

func NewMember(userID string, role Role, joinedAt time.Time) Member {
	return Member{
		userID:   userID,
		role:     role,
		joinedAt: joinedAt,
	}
}

func (m Member) UserID() string {
	return m.userID
}

func (m Member) Role() Role {
	return m.role
}

func (m Member) IsAdmin() bool {
	return m.role == RoleAdmin
}