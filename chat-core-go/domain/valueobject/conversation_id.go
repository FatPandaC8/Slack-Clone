// domain/valueobject/conversation_id.go
package valueobject

import "errors"

// ConversationID represents a conversation identifier
type ConversationID struct {
	value string
}

func NewConversationID(value string) (ConversationID, error) {
	if value == "" {
		return ConversationID{}, errors.New("conversation ID cannot be empty")
	}
	return ConversationID{value: value}, nil
}

func MustConversationID(value string) ConversationID {
	id, err := NewConversationID(value)
	if err != nil {
		panic(err)
	}
	return id
}

func (c ConversationID) Value() string {
	return c.value
}

func (c ConversationID) Equals(other ConversationID) bool {
	return c.value == other.value
}

func (c ConversationID) String() string {
	return c.value
}

func (c ConversationID) IsEmpty() bool {
	return c.value == ""
}