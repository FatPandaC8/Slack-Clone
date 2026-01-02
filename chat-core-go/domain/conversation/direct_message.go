package conversation

import (
	"chat-core-go/domain/message"
	"chat-core-go/domain/user"
)

type DirectMessage struct {
	id       ID
	userA    user.ID
	userB    user.ID
	messages []message.Message
}

func NewDirectMessage(id ID, a, b user.ID) *DirectMessage {
	return &DirectMessage{id: id, userA: a, userB: b}
}

func (d *DirectMessage) ID() ID {
	return d.id
}

func (d *DirectMessage) CanSend(u user.ID) bool {
	return u == d.userA || u == d.userB
}

func (d *DirectMessage) AddMessage(msg message.Message) {
	d.messages = append(d.messages, msg)
}