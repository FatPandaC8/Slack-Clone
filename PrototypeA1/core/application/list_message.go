package application

import (
	"core/port/in"
	"core/port/out"
	"errors"
)

type ListMessagesUseCase struct {
	roomMemberRepo out.RoomMemberRepository
	messageRepo    out.MessageRepository
}

func NewListMessagesUseCase(
	roomMemberRepo out.RoomMemberRepository,
	messageRepo out.MessageRepository,
) *ListMessagesUseCase {
	return &ListMessagesUseCase{
		roomMemberRepo: roomMemberRepo,
		messageRepo:    messageRepo,
	}
}

func (uc *ListMessagesUseCase) ListMessages(
	q in.ListMessagesQuery,
) ([]in.MessageDTO, error) {

	isMember, err := uc.roomMemberRepo.IsMember(q.RoomID, q.UserID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("forbidden")
	}

	messages, err := uc.messageRepo.FindByRoom(q.RoomID)
	if err != nil {
		return nil, err
	}

	result := make([]in.MessageDTO, 0, len(messages))
	for _, m := range messages {
		result = append(result, in.MessageDTO{
			ID:        m.ID(),
			SenderID: m.SenderID(),
			Content:  m.Content(),
			Created:  m.CreatedAt(),
		})
	}

	return result, nil
}
