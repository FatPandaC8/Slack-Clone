package conversation

import (
	"errors"
	"time"

	"chat-core-go/domain/message"
)

type Conversation struct {
	id      	string
	members 	map[string]bool
	messagesIds []string
	createdAt 	time.Time
}

func NewConversation(id string, members []string) (*Conversation, error) {
	if len(members) < 2 {
		return nil, errors.New("conversation must have at least 2 members")
	}

	m := make(map[string]bool)
	for _, u := range members {
		m[u] = true
	}

	return &Conversation{
		id: id,
		members: m,
		messagesIds: []string{},
		createdAt: time.Now(),
	}, nil
}

func (c *Conversation) ID() string {
	return c.id
}

func (c *Conversation) IsMember(u string) bool {
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

func (c *Conversation) MessagesIDs() []string {
	return c.messagesIds
}

func (c *Conversation) MembersList() []string {
	list := make([]string, 0, len(c.members))
	for uid := range c.members {
		list = append(list, uid)
	}
	return list
}