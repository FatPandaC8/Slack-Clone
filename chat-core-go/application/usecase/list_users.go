package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/ports/out"
)

type ListUsers struct {
	users out.UserRepository
}

func NewListUsers(repo out.UserRepository) *ListUsers {
	return &ListUsers{
		users: repo,
	}
}

func (uc *ListUsers) Execute(conversationId string) ([]dto.UserDTO, error) {
	return uc.users.List()
}