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
	users, err := uc.users.List()
	if err != nil {
		return nil, err
	}
	result := make([]dto.UserDTO, 0, len(users))
	for _, u := range users {
		result = append(result, dto.UserDTO{
			ID:   u.ID(),
			Name: u.Name(),
		})
	}
	return result, nil
}

