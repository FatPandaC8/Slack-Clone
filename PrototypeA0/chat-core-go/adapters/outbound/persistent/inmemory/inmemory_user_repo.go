package persistent

import (
	"chat-core-go/domain/user"
	"chat-core-go/domain/valueobject"
	"errors"
	"fmt"
	"sync"
)

// UserRepository is an in-memory implementation of the UserRepository port
type UserRepository struct {
	store   map[string]*user.User // key: userID.Value()
	byEmail map[string]*user.User // key: email.Value()
	mu      sync.RWMutex
	counter int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		store:   make(map[string]*user.User),
		byEmail: make(map[string]*user.User),
		counter: 0,
	}
}

func (r *UserRepository) GenerateID() valueobject.UserID {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.counter++
	return valueobject.MustUserID(fmt.Sprintf("user-%d", r.counter))
}

func (r *UserRepository) Save(u *user.User) error {
	if u == nil {
		return errors.New("user cannot be nil")
	}
	
	r.mu.Lock()
	defer r.mu.Unlock()
	
	// Store by ID and email for quick lookups
	r.store[u.ID().Value()] = u
	r.byEmail[u.Email().Value()] = u
	
	return nil
}

func (r *UserRepository) Load(id valueobject.UserID) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	u, exists := r.store[id.Value()]
	if !exists {
		return nil, errors.New("user not found")
	}
	
	return u, nil
}

func (r *UserRepository) LoadByEmail(email valueobject.Email) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	u, exists := r.byEmail[email.Value()]
	if !exists {
		return nil, errors.New("user not found")
	}
	
	return u, nil
}

func (r *UserRepository) List() ([]*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	users := make([]*user.User, 0, len(r.store))
	for _, u := range r.store {
		users = append(users, u)
	}
	
	return users, nil
}

func (r *UserRepository) Exists(id valueobject.UserID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	_, exists := r.store[id.Value()]
	return exists, nil
}