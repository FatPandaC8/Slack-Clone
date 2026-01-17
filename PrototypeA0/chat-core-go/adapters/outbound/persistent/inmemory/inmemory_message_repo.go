package persistent

import (
	"chat-core-go/domain/message"
	"chat-core-go/domain/valueobject"
	"errors"
	"fmt"
	"sync"
)

// MessageRepository is an in-memory implementation
type MessageRepository struct {
	store   map[string]*message.Message // key: messageID.Value()
	mu      sync.RWMutex
	counter int
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		store:   make(map[string]*message.Message),
		counter: 0,
	}
}

func (r *MessageRepository) GenerateID() valueobject.MessageID {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.counter++
	return valueobject.MustMessageID(fmt.Sprintf("msg-%d", r.counter))
}

func (r *MessageRepository) Save(msg *message.Message) error {
	if msg == nil {
		return errors.New("message cannot be nil")
	}
	
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.store[msg.ID().Value()] = msg
	
	return nil
}

func (r *MessageRepository) Load(id valueobject.MessageID) (*message.Message, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	msg, exists := r.store[id.Value()]
	if !exists {
		return nil, errors.New("message not found")
	}
	
	return msg, nil
}

func (r *MessageRepository) LoadByConversation(convID valueobject.ConversationID) ([]*message.Message, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	result := make([]*message.Message, 0)
	
	for _, msg := range r.store {
		if msg.ConversationID().Equals(convID) {
			result = append(result, msg)
		}
	}
	
	return result, nil
}

func (r *MessageRepository) Delete(id valueobject.MessageID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if _, exists := r.store[id.Value()]; !exists {
		return errors.New("message not found")
	}
	
	delete(r.store, id.Value())
	return nil
}