package usecase

import (
	"errors"

	"chat-core-go/application/dto"
	"chat-core-go/domain/message"
	"chat-core-go/ports/out"
)

type SendMessage struct {
	conversations  	out.ConversationRepository
	messages		out.MessageRepository
	publisher 		out.MessagePublisher
}

func NewSendMessage(
	c out.ConversationRepository,
	m out.MessageRepository,
	p out.MessagePublisher,
) *SendMessage {
	return &SendMessage{c, m, p}
}

func (uc *SendMessage) Execute(cmd dto.SendMessageCommand) error {
	conv, err := uc.conversations.Load(cmd.ConversationID)
	if err != nil {
		return err
	}

	if !conv.IsMember(cmd.SenderID) {
		return errors.New("permission denied")
	}

	msg := message.NewMessage(
		message.NewID(cmd.MessageID),
		cmd.SenderID,
		cmd.ConversationID,
		message.NewContent(cmd.Text),
	)

	conv.AddMessage(msg)
	uc.messages.Save(msg)
	uc.publisher.Publish(msg)

	return nil
}