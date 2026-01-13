package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/user"
	"chat-core-go/ports/out"
	"chat-core-go/ports/service"
	"errors"
)

type RegisterUser struct {
	users   out.UserRepository
	hasher  service.PasswordHasher
}

func NewRegisterUser(users out.UserRepository, hasher service.PasswordHasher) *RegisterUser {
	return &RegisterUser{users, hasher}
}

func (uc *RegisterUser) Execute(cmd dto.RegisterUserCommand) (*dto.UserDTO, error) {
	existing, _ := uc.users.LoadByEmail(cmd.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hash, err := uc.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, err
	}

	id := uc.users.GenerateID()
	u, err := user.NewUser(id, cmd.Name, cmd.Email, hash)
	if err != nil {
		return nil, err
	}

	err = uc.users.Save(u)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID: u.ID(),
		Name: u.Name(),
	}, nil
}