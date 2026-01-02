package persistent

import "chat-core-go/domain/message"

type InMemoryMessageRepo struct {
	messages []message.Message
}

func NewInMemoryMessageRepo() *InMemoryMessageRepo {
	return &InMemoryMessageRepo{}
}

func (r *InMemoryMessageRepo) Save(msg message.Message) error {
	r.messages = append(r.messages, msg)
	return nil
}