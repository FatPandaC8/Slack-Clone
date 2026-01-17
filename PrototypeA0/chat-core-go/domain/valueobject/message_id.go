// domain/valueobject/message_id.go
package valueobject

import "errors"

// MessageID represents a message identifier
type MessageID struct {
	value string
}

func NewMessageID(value string) (MessageID, error) {
	if value == "" {
		return MessageID{}, errors.New("message ID cannot be empty")
	}
	return MessageID{value: value}, nil
}

func MustMessageID(value string) MessageID {
	id, err := NewMessageID(value)
	if err != nil {
		panic(err)
	}
	return id
}

func (m MessageID) Value() string {
	return m.value
}

func (m MessageID) Equals(other MessageID) bool {
	return m.value == other.value
}

func (m MessageID) String() string {
	return m.value
}

func (m MessageID) IsEmpty() bool {
	return m.value == ""
}