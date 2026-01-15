// domain/valueobject/message_content.go
package valueobject

import (
	"errors"
	"strings"
)

// MessageContent represents the text content of a message
// Domain rules: Non-empty, max 5000 characters, trimmed
type MessageContent struct {
	value string
}

const (
	MinMessageLength = 1
	MaxMessageLength = 5000
)

func NewMessageContent(value string) (MessageContent, error) {
	trimmed := strings.TrimSpace(value)
	
	if len(trimmed) < MinMessageLength {
		return MessageContent{}, errors.New("message content cannot be empty")
	}
	if len(trimmed) > MaxMessageLength {
		return MessageContent{}, errors.New("message content exceeds maximum length")
	}
	
	return MessageContent{value: trimmed}, nil
}

func MustMessageContent(value string) MessageContent {
	content, err := NewMessageContent(value)
	if err != nil {
		panic(err)
	}
	return content
}

func (m MessageContent) Value() string {
	return m.value
}

func (m MessageContent) Length() int {
	return len(m.value)
}

func (m MessageContent) IsEmpty() bool {
	return m.value == ""
}

func (m MessageContent) String() string {
	return m.value
}