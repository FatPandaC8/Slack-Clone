package out

import (
	"chat-core-go/domain/user"
)

type UserRepository interface {
	Load(id string) (*user.User, error) // load by user id
	LoadByEmail(email string) (*user.User, error) // load by email
	Save(u *user.User) error // Save user
	GenerateID() string // Generate user their id
	List() ([]*user.User, error) // list the users by their name and id
}