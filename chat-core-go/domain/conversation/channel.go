package conversation

import (
	"chat-core-go/domain/message"
	"chat-core-go/domain/user"
)

type Channel struct {
	id 			ID
	members 	map[user.ID]bool
	messages 	[]message.Message
}

func NewChannel(id ID, members []user.ID) *Channel {
	m := make(map[user.ID]bool)
	for _, uid := range members {
		m[uid] = true
	}
	return &Channel{id: id, members: m}
}

func (c *Channel) ID() ID {
	return c.id
}

func (c *Channel) CanSend(u user.ID) bool {
	return c.members[u]
}

func (c *Channel) AddMessage(msg message.Message) {
	c.messages = append(c.messages, msg)
}