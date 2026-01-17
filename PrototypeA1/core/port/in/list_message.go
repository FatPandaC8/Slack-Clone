package in

import "time"

type MessageDTO struct {
	ID        string
	SenderID string
	Content  string
	Created  time.Time
}

type ListMessagesQuery struct {
	UserID string
	RoomID string
}

type ListMessagesUseCase interface {
	ListMessages(q ListMessagesQuery) ([]MessageDTO, error)
}