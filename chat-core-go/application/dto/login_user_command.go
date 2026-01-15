package dto

import "chat-core-go/domain/valueobject"

type LoginUserCommand struct {
	Email    valueobject.Email
	Password string // Plaintext - will be compared with hash
}

type LoginUserResult struct {
	UserID valueobject.UserID
	Name   valueobject.UserName
	Email  valueobject.Email
}