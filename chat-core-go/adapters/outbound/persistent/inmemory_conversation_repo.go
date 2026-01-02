package persistent

import (
	"errors"

	"chat-core-go/domain/conversation"
)

type InMemoryConversationRepo struct {
	store map[string]conversation.Conversation
}

func NewInMemoryConversationRepo() *InMemoryConversationRepo {
	repo := &InMemoryConversationRepo{
		store: make(map[string]conversation.Conversation),
	}

	return repo
}

func (r *InMemoryConversationRepo) Load(id string) (conversation.Conversation, error) {
	conv, ok := r.store[id]
	if !ok {
		return nil, errors.New("conversation not found")
	}
	return conv, nil
}

func (r *InMemoryConversationRepo) Save(conv conversation.Conversation) error {
	r.store[string(conv.ID())] = conv
	return nil
}
