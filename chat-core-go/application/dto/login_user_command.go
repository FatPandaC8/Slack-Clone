package dto

type LoginUserCommand struct {
	Email    string
	Password string
}

type LoginResultDTO struct {
	UserID string
	Name   string
	Email  string
	Token  string
}