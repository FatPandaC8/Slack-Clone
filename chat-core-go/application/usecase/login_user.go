package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/ports/out"
	"chat-core-go/ports/service"
	"errors"
)

// LoginUser is a use case for authenticating users
type LoginUser struct {
	users  out.UserRepository
	hasher service.PasswordHasher
}

func NewLoginUser(
	users out.UserRepository,
	hasher service.PasswordHasher,
) *LoginUser {
	return &LoginUser{
		users:  users,
		hasher: hasher,
	}
}

// Execute performs the login user use case
func (uc *LoginUser) Execute(cmd dto.LoginUserCommand) (*dto.LoginUserResult, error) {
	// 1. Validate input
	if cmd.Email.IsEmpty() {
		return nil, errors.New("email is required")
	}
	if cmd.Password == "" {
		return nil, errors.New("password is required")
	}
	
	// 2. Load user by email
	user, err := uc.users.LoadByEmail(cmd.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	
	// 3. Verify password
	err = uc.hasher.Compare(user.PasswordHash(), cmd.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	
	// 4. Return result
	return &dto.LoginUserResult{
		UserID: user.ID(),
		Name:   user.Name(),
		Email:  user.Email(),
	}, nil
}