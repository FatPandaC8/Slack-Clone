// domain/valueobject/invite_code.go
package valueobject

import (
	"errors"
	"regexp"
)

// InviteCode represents a conversation invite code
// Domain rule: Must be 6 uppercase alphanumeric characters
type InviteCode struct {
	value string
}

var inviteCodePattern = regexp.MustCompile(`^[A-Z0-9]{6}$`)

func NewInviteCode(value string) (InviteCode, error) {
	if value == "" {
		return InviteCode{}, errors.New("invite code cannot be empty")
	}
	if !inviteCodePattern.MatchString(value) {
		return InviteCode{}, errors.New("invite code must be 6 uppercase alphanumeric characters")
	}
	return InviteCode{value: value}, nil
}

func MustInviteCode(value string) InviteCode {
	code, err := NewInviteCode(value)
	if err != nil {
		panic(err)
	}
	return code
}

func (i InviteCode) Value() string {
	return i.value
}

func (i InviteCode) Equals(other InviteCode) bool {
	return i.value == other.value
}

func (i InviteCode) String() string {
	return i.value
}

func (i InviteCode) IsEmpty() bool {
	return i.value == ""
}