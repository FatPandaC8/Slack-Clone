package persistent

import (
	"errors"

	"chat-core-go/domain/conversation"
	"chat-core-go/domain/user"
)

type InMemoryConversationRepo struct {
	store map[string]*conversation.Conversation
}

func NewInMemoryConversationRepo() *InMemoryConversationRepo {
	repo := &InMemoryConversationRepo{
		store: make(map[string]*conversation.Conversation),
	}

	return repo
}

func (r *InMemoryConversationRepo) Load(id string) (*conversation.Conversation, error) {
	conv, ok := r.store[id]
	if !ok {
		return nil, errors.New("conversation not found")
	}
	return conv, nil
}

func (r *InMemoryConversationRepo) Save(conv *conversation.Conversation) error {
	r.store[string(conv.ID())] = conv
	return nil
}

func (r *InMemoryConversationRepo) FindByMember(uid user.ID) ([]conversation.Conversation, error) {
	var result []conversation.Conversation
	for _, conv := range r.store {
		if conv.IsMember(uid) {
			result = append(result, *conv)
		}
	}
	return result, nil
}