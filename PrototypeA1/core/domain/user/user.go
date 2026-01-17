package user

import "time"

type User struct {
	userID 	string
	name 	string
	email 	string
	createdAt time.Time
}

func NewUser(id, name, email string, createdAt time.Time) *User {
	return &User{
		userID:        id,
		name:      name,
		email:     email,
		createdAt: createdAt,
	}
}

func (u *User) ID() string {
	return u.userID
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}