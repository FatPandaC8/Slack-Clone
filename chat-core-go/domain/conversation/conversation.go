package conversation

import (
	"errors"
	"time"

	"chat-core-go/domain/message"
	"chat-core-go/domain/user"
)

type Conversation struct {
	id      	ID
	members 	map[user.ID]bool
	messagesIds []message.ID
	createdAt 	time.Time
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
		messagesIds: []message.ID{},
		createdAt: time.Now(),
	}, nil
}

func (c *Conversation) ID() ID {
	return c.id
}

func (c *Conversation) IsMember(u user.ID) bool {
	return c.members[u]
}

func (c *Conversation) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Conversation) AddMessage(msg message.Message) error {
	if !c.IsMember(msg.Sender()) {
		return errors.New("sender is not a member")
	}
	c.messagesIds = append(c.messagesIds, msg.ID())
	return nil
}

func (c *Conversation) MessagesIDs() []message.ID {
	return c.messagesIds
}

func (c *Conversation) MembersList() []user.ID {
	list := make([]user.ID, 0, len(c.members))
	for uid := range c.members {
		list = append(list, uid)
	}
	return list
}