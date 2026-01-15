// application/usecase/register_user.go
package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/user"
	"chat-core-go/ports/out"
	"chat-core-go/ports/service"
	"errors"
)

// RegisterUser is a use case for registering new users
type RegisterUser struct {
	users  out.UserRepository
	hasher service.PasswordHasher
}

func NewRegisterUser(
	users out.UserRepository,
	hasher service.PasswordHasher,
) *RegisterUser {
	return &RegisterUser{
		users:  users,
		hasher: hasher,
	}
}

// Execute performs the register user use case
func (uc *RegisterUser) Execute(cmd dto.RegisterUserCommand) (*dto.RegisterUserResult, error) {
	// 1. Validate input
	if cmd.Name.IsEmpty() {
		return nil, errors.New("name is required")
	}
	if cmd.Email.IsEmpty() {
		return nil, errors.New("email is required")
	}
	if cmd.Password == "" {
		return nil, errors.New("password is required")
	}
	if len(cmd.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}
	
	// 2. Check if email already exists
	existing, _ := uc.users.LoadByEmail(cmd.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}
	
	// 3. Hash password
	passwordHash, err := uc.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	
	// 4. Create user entity
	userID := uc.users.GenerateID()
	newUser, err := user.NewUser(
		userID,
		cmd.Name,
		cmd.Email,
		passwordHash,
	)
	if err != nil {
		return nil, err
	}
	
	// 5. Persist
	err = uc.users.Save(newUser)
	if err != nil {
		return nil, err
	}
	
	// 6. Return result
	return &dto.RegisterUserResult{
		UserID: userID,
		Name:   cmd.Name,
		Email:  cmd.Email,
	}, nil
}