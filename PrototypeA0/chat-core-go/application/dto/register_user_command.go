package dto

import "chat-core-go/domain/valueobject"

// RegisterUserCommand contains input for registration
type RegisterUserCommand struct {
	Name     valueobject.UserName
	Email    valueobject.Email
	Password string // Plaintext - will be hashed
}

// RegisterUserResult contains output from registration
type RegisterUserResult struct {
	UserID valueobject.UserID
	Name   valueobject.UserName
	Email  valueobject.Email
}