package user

import "time"

type User struct {
	id 				string
	name 			string
	email 			string
	passwordHash 	string
	createdAt 		time.Time
}

func NewUser(id string, name, email, passwordHash string) (*User, error) {
	// let the client filter the invalid name, email and password Hash
	
	return &User{
		id: id,
		name: name,
		email: email,
		passwordHash: passwordHash,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}