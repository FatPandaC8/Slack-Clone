package conversation

import (
	"chat-core-go/domain/valueobject"
	"errors"
	"time"
)

// Conversation is an AGGREGATE ROOT
// It controls access to messages and enforces invariants
type Conversation struct {
	id         valueobject.ConversationID
	name       valueobject.ConversationName
	inviteCode valueobject.InviteCode
	members    []valueobject.UserID
	messageIDs []valueobject.MessageID
	createdAt  time.Time
}

// NewConversation creates a new conversation (factory method)
func NewConversation(
	id valueobject.ConversationID,
	name valueobject.ConversationName,
	inviteCode valueobject.InviteCode,
	creatorID valueobject.UserID,
) (*Conversation, error) {
	if id.IsEmpty() {
		return nil, errors.New("conversation ID is required")
	}
	if name.IsEmpty() {
		return nil, errors.New("conversation name is required")
	}
	if inviteCode.IsEmpty() {
		return nil, errors.New("invite code is required")
	}
	if creatorID.IsEmpty() {
		return nil, errors.New("creator ID is required")
	}
	
	return &Conversation{
		id:         id,
		name:       name,
		inviteCode: inviteCode,
		members:    []valueobject.UserID{creatorID},
		messageIDs: []valueobject.MessageID{},
		createdAt:  time.Now(),
	}, nil
}

// ReconstructConversation recreates from persistence
func ReconstructConversation(
	id valueobject.ConversationID,
	name valueobject.ConversationName,
	inviteCode valueobject.InviteCode,
	members []valueobject.UserID,
	messageIDs []valueobject.MessageID,
	createdAt time.Time,
) *Conversation {
	return &Conversation{
		id:         id,
		name:       name,
		inviteCode: inviteCode,
		members:    members,
		messageIDs: messageIDs,
		createdAt:  createdAt,
	}
}

// Getters
func (c *Conversation) ID() valueobject.ConversationID {
	return c.id
}

func (c *Conversation) Name() valueobject.ConversationName {
	return c.name
}

func (c *Conversation) InviteCode() valueobject.InviteCode {
	return c.inviteCode
}

func (c *Conversation) Members() []valueobject.UserID {
	// Return copy to prevent external modification
	membersCopy := make([]valueobject.UserID, len(c.members))
	copy(membersCopy, c.members)
	return membersCopy
}

func (c *Conversation) MessageIDs() []valueobject.MessageID {
	// Return copy
	msgsCopy := make([]valueobject.MessageID, len(c.messageIDs))
	copy(msgsCopy, c.messageIDs)
	return msgsCopy
}

func (c *Conversation) CreatedAt() time.Time {
	return c.createdAt
}

// Domain behaviors (business logic)

// AddMember adds a user to the conversation (idempotent)
func (c *Conversation) AddMember(userID valueobject.UserID) error {
	if userID.IsEmpty() {
		return errors.New("user ID cannot be empty")
	}
	
	// Check if already a member (idempotent)
	if c.HasMember(userID) {
		return nil
	}
	
	c.members = append(c.members, userID)
	return nil
}

// RemoveMember removes a user from the conversation
func (c *Conversation) RemoveMember(userID valueobject.UserID) error {
	if userID.IsEmpty() {
		return errors.New("user ID cannot be empty")
	}
	
	// Cannot remove if it's the last member
	if len(c.members) == 1 && c.HasMember(userID) {
		return errors.New("cannot remove the last member")
	}
	
	for i, member := range c.members {
		if member.Equals(userID) {
			// Remove by swapping with last and truncating
			c.members[i] = c.members[len(c.members)-1]
			c.members = c.members[:len(c.members)-1]
			return nil
		}
	}
	
	return errors.New("user is not a member")
}

// HasMember checks if a user is a member
func (c *Conversation) HasMember(userID valueobject.UserID) bool {
	for _, member := range c.members {
		if member.Equals(userID) {
			return true
		}
	}
	return false
}

// AddMessage adds a message to the conversation (idempotent)
func (c *Conversation) AddMessage(messageID valueobject.MessageID) error {
	if messageID.IsEmpty() {
		return errors.New("message ID cannot be empty")
	}
	
	// Check if already exists (idempotent)
	for _, id := range c.messageIDs {
		if id.Equals(messageID) {
			return nil
		}
	}
	
	c.messageIDs = append(c.messageIDs, messageID)
	return nil
}

// MemberCount returns the number of members
func (c *Conversation) MemberCount() int {
	return len(c.members)
}

// MessageCount returns the number of messages
func (c *Conversation) MessageCount() int {
	return len(c.messageIDs)
}

// Rename changes the conversation name
func (c *Conversation) Rename(newName valueobject.ConversationName) error {
	if newName.IsEmpty() {
		return errors.New("name cannot be empty")
	}
	c.name = newName
	return nil
}