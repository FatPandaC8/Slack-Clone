package out

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/user"
)

type UserRepository interface {
	Load(id string) (*user.User, error)
	Save(u *user.User) error 
	GenerateID() (string)
	List() ([]dto.UserDTO, error)
}