package persistent

import (
	"errors"
	"fmt"
	"sync"

	"chat-core-go/domain/conversation"
)

type InMemoryConversationRepo struct {
	store map[string]*conversation.Conversation
	mu sync.Mutex
	counter int
}

func NewInMemoryConversationRepo() *InMemoryConversationRepo {
	return &InMemoryConversationRepo{
		store: make(map[string]*conversation.Conversation),
	}
}

func (r *InMemoryConversationRepo) GenerateID() string {
	r.mu.Lock()

	defer r.mu.Unlock()
	r.counter++
	return fmt.Sprintf("conv-%d", r.counter)
}

func (r *InMemoryConversationRepo) Load(id string) (*conversation.Conversation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	conv, ok := r.store[id]
	if !ok {
		return nil, errors.New("conversation not found")
	}
	return conv, nil
}

func (r *InMemoryConversationRepo) Save(conv *conversation.Conversation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if conv == nil {
		return errors.New("conversation is null")
	}
	r.store[conv.ID()] = conv
	return nil
}

func (r *InMemoryConversationRepo) FindByMember(uid string) ([]conversation.Conversation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []conversation.Conversation
	for _, conv := range r.store {
		if conv.IsMember(uid) {
			result = append(result, *conv)
		}
	}
	return result, nil
}