package persistent

import (
	"chat-core-go/domain/message"
	"fmt"
	"sync"
)

type InMemoryMessageRepo struct {
	messages []message.Message
	mu sync.Mutex
	counter int
}

func NewInMemoryMessageRepo() *InMemoryMessageRepo {
	return &InMemoryMessageRepo{
		messages: []message.Message{},
	}
}

func (r *InMemoryMessageRepo) GenerateID() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter++
	return fmt.Sprintf("msg-%d", r.counter)
}

func (r *InMemoryMessageRepo) Save(msg message.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.messages = append(r.messages, msg)
	return nil
}

func (r *InMemoryMessageRepo) Load(id string) (message.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, msg := range r.messages {
		if string(msg.ID()) == id {
			return msg, nil
		}
	}

	return message.Message{}, fmt.Errorf("message not found")
}