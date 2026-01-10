package conversation

import (
	"time"
)

type Conversation struct {
	id      	string
	name 		string
	inviteCode  string
	members 	[]string
	messagesIds []string
	createdAt 	time.Time // for other kind of plugins if i like
}

func NewConversation(id, name, inviteCode, creatorID string) (*Conversation, error) {
	return &Conversation{
		id: id,
		name: name,
		inviteCode: inviteCode,
		members: []string{creatorID},
		messagesIds: []string{},
		createdAt: time.Now(),
	}, nil
}

func (c *Conversation) ID() string            { return c.id }
func (c *Conversation) Name() string          { return c.name }
func (c *Conversation) InviteCode() string    { return c.inviteCode }
func (c *Conversation) Members() []string     { return c.members }
func (c *Conversation) MessageIDs() []string  { return c.messagesIds }

func (c *Conversation) AddMember(uid string) {
    // avoid duplicates
    for _, m := range c.members {
        if m == uid {
            return
        }
    }
    c.members = append(c.members, uid)
}

func (c *Conversation) HasMember(uid string) bool {
	for _, m := range c.members {
		if m == uid {
			return true
		}
	}
	return false
}

func (c *Conversation) AddMessageID(mid string) {
	// avoid duplicates just in case
	for _, id := range c.messagesIds {
		if id == mid {
			return
		}
	}

	c.messagesIds = append(c.messagesIds, mid)
}
