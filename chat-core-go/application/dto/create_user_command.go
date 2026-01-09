package dto

type CreateUserCommand struct {
	Name 		string
	Email 		string
	PasswordHash string
}