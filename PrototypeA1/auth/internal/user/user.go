package user

import "time"

type User struct {
	userID 		string
	name 		string
	email		string
	password 	string
	createdAt 	time.Time
}

func NewUser(
	userID, name, email, password string,
	createdAt time.Time,
) *User {
	return &User{
		userID:    userID,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
	}
}

func (u User) ID() string {
	return u.userID
}

func (u User) Password() string {
	return u.password
}

type UserRepository interface {
	Create(name, email, password string) (string, error)
	Save(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(userID string) (*User, error) 
}