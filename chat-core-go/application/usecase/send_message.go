// application/usecase/send_message.go
package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/domain/message"
	"chat-core-go/ports/out"
	"errors"
)

// SendMessage is a use case for sending messages in a conversation
type SendMessage struct {
	conversations out.ConversationRepository
	messages      out.MessageRepository
	users         out.UserRepository
}

func NewSendMessage(
	conversations out.ConversationRepository,
	messages out.MessageRepository,
	users out.UserRepository,
) *SendMessage {
	return &SendMessage{
		conversations: conversations,
		messages:      messages,
		users:         users,
	}
}

// Execute performs the send message use case
func (uc *SendMessage) Execute(cmd dto.SendMessageCommand) error {
	// 1. Validate principal
	if cmd.Principal == nil {
		return errors.New("principal is required")
	}
	if cmd.Principal.IsExpired() {
		return errors.New("principal expired")
	}
	
	// 2. Load conversation (aggregate root)
	conv, err := uc.conversations.Load(cmd.ConversationID)
	if err != nil {
		return errors.New("conversation not found")
	}
	
	// 3. Authorization: Check membership (domain rule)
	if !conv.HasMember(cmd.Principal.UserID()) {
		return errors.New("not authorized: user is not a member of this conversation")
	}
	
	// 4. Create message entity
	messageID := uc.messages.GenerateID()
	msg := message.NewMessage(
		messageID,
		cmd.Principal.UserID(),
		cmd.ConversationID,
		cmd.Content,
	)
	
	// 5. Add message to conversation (aggregate invariant)
	err = conv.AddMessage(messageID)
	if err != nil {
		return err
	}
	
	// 6. Persist changes (Unit of Work pattern - could be transactional)
	err = uc.messages.Save(msg)
	if err != nil {
		return err
	}
	
	err = uc.conversations.Save(conv)
	if err != nil {
		return err
	}
	
	return nil
}