package persistent

import (
	"chat-core-go/domain/user"
	"errors"
	"fmt"
	"sync"
)

type InMemoryUserRepo struct {
	store   map[string]*user.User
	mu      sync.Mutex
	counter int
}

func NewInmemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		store: make(map[string]*user.User),
	}
}

func (r *InMemoryUserRepo) GenerateID() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter++
	return string(fmt.Sprintf("user-%d", r.counter))
}

func (r *InMemoryUserRepo) Save(user *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user == nil {
		return errors.New("user is null")
	}

	r.store[user.ID()] = user
	return nil
}

func (r *InMemoryUserRepo) Load(id string) (*user.User, error) {
	r.mu.Lock()

	defer r.mu.Unlock()

	u, ok := r.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *InMemoryUserRepo) List() ([]*user.User, error) {
	users := make([]*user.User, 0, len(r.store))
	for _, u := range r.store {
		users = append(users, u)
	}
	return users, nil
}
