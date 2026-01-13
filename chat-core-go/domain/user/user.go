package user

import (
	"errors"
	"time"
)

type User struct {
	id 				string
	name 			string
	email 			string
	passwordHash 	string
	createdAt 		time.Time
}

func NewUser(id string, name, email, passwordHash string) (*User, error) {
	if name == "" || email == "" || passwordHash == "" {
		return nil, errors.New("invalid user fields")
	}
	
	return &User{
		id: id,
		name: name,
		email: email,
		passwordHash: passwordHash,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ID() string           { return u.id }
func (u *User) Name() string         { return u.name }
func (u *User) Email() string        { return u.email }
func (u *User) PasswordHash() string { return u.passwordHash }