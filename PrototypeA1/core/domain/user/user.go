package user

import (
	valueobject "core/domain/valueobject/user"
	"time"
)

type User struct {
	userID 	valueobject.UserID
	name 	string
	email 	string
	createdAt time.Time
}

func NewUser(id valueobject.UserID, name, email string, createdAt time.Time) *User {
	return &User{
		userID:        id,
		name:      name,
		email:     email,
		createdAt: createdAt,
	}
}

func (u *User) ID() string {
	return u.userID.String()
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}