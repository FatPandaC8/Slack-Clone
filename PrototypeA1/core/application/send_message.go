package application

import (
	"core/domain/message"
	"core/port/in"
	"core/port/out"
	"errors"
	"time"

	"github.com/google/uuid"
)

type SendMessageUseCase struct {
	roomMemberRepo out.RoomMemberRepository
	messageRepo    out.MessageRepository
	broadcaster    out.MessageBroadcaster
}

func NewSendMessageUseCase(
	roomMemberRepo out.RoomMemberRepository,
	messageRepo out.MessageRepository,
	broadcaster out.MessageBroadcaster,
) *SendMessageUseCase {
	return &SendMessageUseCase{
		roomMemberRepo: roomMemberRepo,
		messageRepo:    messageRepo,
		broadcaster:    broadcaster,
	}
}

func (uc *SendMessageUseCase) SendMessage(cmd in.SendMessageCommand) error {
	isMember, err := uc.roomMemberRepo.IsMember(cmd.RoomID, cmd.UserID)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("forbidden")
	}

	msg := message.NewMessage(
		uuid.NewString(),
		cmd.RoomID,
		cmd.UserID,
		cmd.Content,
		time.Now(),
	)

	if err := uc.messageRepo.Save(msg); err != nil {
		return err
	}

	return uc.broadcaster.Broadcast(cmd.RoomID, msg)
}