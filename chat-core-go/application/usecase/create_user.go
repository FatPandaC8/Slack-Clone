package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/user"
	"chat-core-go/ports/out"
)

type CreateUser struct {
	users out.UserRepository
}

func NewCreateUser (
	users out.UserRepository,
) *CreateUser {
	return &CreateUser{
		users: users,
	}
}

func (uc *CreateUser) Execute(cmd dto.CreateUserCommand) (*dto.UserDTO, error) {
	id := uc.users.GenerateID()
	user, err := user.NewUser(
		id, 
		cmd.Name,
		cmd.Email,
		cmd.PasswordHash,
	)
	if err != nil {
		return nil, err
	}

	err = uc.users.Save(user)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID: user.ID(),
		Name: user.Name(),
	}, nil
}