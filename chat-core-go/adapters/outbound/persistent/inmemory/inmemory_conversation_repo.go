package persistent

import (
	"chat-core-go/domain/conversation"
	"chat-core-go/domain/valueobject"
	"errors"
	"fmt"
	"sync"
)

// ConversationRepository is an in-memory implementation
type ConversationRepository struct {
	store    map[string]*conversation.Conversation // key: conversationID.Value()
	byCode   map[string]*conversation.Conversation // key: inviteCode.Value()
	mu       sync.RWMutex
	counter  int
}

func NewConversationRepository() *ConversationRepository {
	return &ConversationRepository{
		store:   make(map[string]*conversation.Conversation),
		byCode:  make(map[string]*conversation.Conversation),
		counter: 0,
	}
}

func (r *ConversationRepository) GenerateID() valueobject.ConversationID {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.counter++
	return valueobject.MustConversationID(fmt.Sprintf("conv-%d", r.counter))
}

func (r *ConversationRepository) Save(conv *conversation.Conversation) error {
	if conv == nil {
		return errors.New("conversation cannot be nil")
	}
	
	r.mu.Lock()
	defer r.mu.Unlock()
	
	// Store by ID and invite code
	r.store[conv.ID().Value()] = conv
	r.byCode[conv.InviteCode().Value()] = conv
	
	return nil
}

func (r *ConversationRepository) Load(id valueobject.ConversationID) (*conversation.Conversation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	conv, exists := r.store[id.Value()]
	if !exists {
		return nil, errors.New("conversation not found")
	}
	
	return conv, nil
}

func (r *ConversationRepository) LoadByInviteCode(code valueobject.InviteCode) (*conversation.Conversation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	conv, exists := r.byCode[code.Value()]
	if !exists {
		return nil, errors.New("conversation not found")
	}
	
	return conv, nil
}

func (r *ConversationRepository) FindByMember(userID valueobject.UserID) ([]*conversation.Conversation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	result := make([]*conversation.Conversation, 0)
	
	for _, conv := range r.store {
		if conv.HasMember(userID) {
			result = append(result, conv)
		}
	}
	
	return result, nil
}

func (r *ConversationRepository) Exists(id valueobject.ConversationID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	_, exists := r.store[id.Value()]
	return exists, nil
}