// domain/valueobject/conversation_name.go
package valueobject

import (
	"errors"
	"strings"
)

// ConversationName represents a conversation's display name
type ConversationName struct {
	value string
}

const (
	MinConversationNameLength = 1
	MaxConversationNameLength = 200
)

func NewConversationName(value string) (ConversationName, error) {
	trimmed := strings.TrimSpace(value)
	
	if len(trimmed) < MinConversationNameLength {
		return ConversationName{}, errors.New("conversation name cannot be empty")
	}
	if len(trimmed) > MaxConversationNameLength {
		return ConversationName{}, errors.New("conversation name too long")
	}
	
	return ConversationName{value: trimmed}, nil
}

func MustConversationName(value string) ConversationName {
	name, err := NewConversationName(value)
	if err != nil {
		panic(err)
	}
	return name
}

func (c ConversationName) Value() string {
	return c.value
}

func (c ConversationName) String() string {
	return c.value
}

func (c ConversationName) IsEmpty() bool {
	return c.value == ""
}