package usecase

import (
	"chat-core-go/application/dto"
	"chat-core-go/ports/out"
	"time"
)

type GetConversation struct {
	conversations out.ConversationRepository
	messages out.MessageRepository
	users out.UserRepository
}

func NewGetConversation(
	convs out.ConversationRepository,
	msgs out.MessageRepository,
	users out.UserRepository,
) *GetConversation {
	return &GetConversation{
		conversations: convs,
		messages:      msgs,
		users:         users,
	}
}

func (uc *GetConversation) Execute(id string) (*dto.GetConversationDTO, error) {
	conv, err := uc.conversations.Load(id)
	if err != nil {
		return nil, err
	}

	var messages []dto.MessageDTO
	for _, msgID := range conv.MessageIDs() {
		msg, err := uc.messages.Load(msgID)
		if err != nil {
			continue // skip error messages
		}
		messages = append(messages, dto.MessageDTO{
			ID: msg.ID(),
			SenderID:  msg.Sender(),
			Content:   msg.Content().Value(),
			CreatedAt: msg.CreatedAt().Format(time.RFC3339),
		})
	}

	var members []dto.UserDTO
	for _, uid := range conv.Members() {
		userObj, err := uc.users.Load(uid)
		if err != nil {
			members = append(members, dto.UserDTO{ID: uid})
			continue
		}
		members = append(members, dto.UserDTO{
			ID:   uid,
			Name: userObj.Name(),
		})
	}

	return &dto.GetConversationDTO{
		ID:       conv.ID(),
		Members:  members,
		Messages: messages,
	}, nil
}