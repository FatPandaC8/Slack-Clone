package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/ports/out"
	"chat-core-go/ports/service"
	"errors"
)

type LoginUser struct {
	users   out.UserRepository
	hasher  service.PasswordHasher
	tokens  service.TokenService
}

func NewLoginUser(users out.UserRepository, hasher service.PasswordHasher, tokens service.TokenService) *LoginUser {
	return &LoginUser{users, hasher, tokens}
}

func (uc *LoginUser) Execute(cmd dto.LoginUserCommand) (*dto.LoginResultDTO, error) {
	u, err := uc.users.LoadByEmail(cmd.Email)
	if err != nil || u == nil {
		return nil, errors.New("invalid email or password")
	}

	err = uc.hasher.Compare(u.PasswordHash(), cmd.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := uc.tokens.GenerateToken(u.ID())
	if err != nil {
		return nil, err
	}

	return &dto.LoginResultDTO{
		UserID: u.ID(),
		Name:   u.Name(),
		Email: u.Email(),
		Token:  token,
	}, nil
}