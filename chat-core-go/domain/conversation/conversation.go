package conversation

import (
	"errors"

	"chat-core-go/domain/message"
	"chat-core-go/domain/user"
)

type Conversation struct {
	id      ID
	members map[user.ID]bool
	messages []message.Message
}

func NewConversation(id ID, members []user.ID) (*Conversation, error) {
	if len(members) < 2 {
		return nil, errors.New("conversation must have at least 2 members")
	}

	m := make(map[user.ID]bool)
	for _, u := range members {
		m[u] = true
	}

	return &Conversation{
		id: id,
		members: m,
		messages: []message.Message{},
	}, nil
}

func (c *Conversation) ID() ID {
	return c.id
}

func (c *Conversation) IsMember(u user.ID) bool {
	return c.members[u]
}

func (c *Conversation) AddMessage(msg message.Message) error {
	if !c.IsMember(msg.Sender()) {
		return errors.New("sender is not a member")
	}
	c.messages = append(c.messages, msg)
	return nil
}

func (c *Conversation) Messages() []message.Message {
	return c.messages
}
