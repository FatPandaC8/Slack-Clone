package persistent

import (
	"errors"
	"fmt"
	"sync"

	"chat-core-go/domain/conversation"
)

type InMemoryConversationRepo struct {
	storeID map[string]*conversation.Conversation
	storeCode map[string]*conversation.Conversation
	mu sync.Mutex
	counter int
}

func NewInMemoryConversationRepo() *InMemoryConversationRepo {
	return &InMemoryConversationRepo{
		storeID: make(map[string]*conversation.Conversation),
		storeCode: make(map[string]*conversation.Conversation),
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

	conv, ok := r.storeID[id]
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
	r.storeID[conv.ID()] = conv
	r.storeCode[conv.InviteCode()] = conv
	return nil
}

func (r *InMemoryConversationRepo) FindByMember(uid string) ([]conversation.Conversation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []conversation.Conversation
	for _, conv := range r.storeID {
		if conv.HasMember(uid) {
			result = append(result, *conv)
		}
	}
	return result, nil
}

func (r *InMemoryConversationRepo) FindByInviteCode(code string) (*conversation.Conversation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	conv, ok := r.storeCode[code]
	if !ok {
		return nil, errors.New("conversation not found")
	}
	return conv, nil
}