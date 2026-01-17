// domain/valueobject/user_name.go
package valueobject

import (
	"errors"
	"strings"
)

// UserName represents a user's display name
type UserName struct {
	value string
}

const (
	MinNameLength = 1
	MaxNameLength = 100
)

func NewUserName(value string) (UserName, error) {
	trimmed := strings.TrimSpace(value)
	
	if len(trimmed) < MinNameLength {
		return UserName{}, errors.New("user name cannot be empty")
	}
	if len(trimmed) > MaxNameLength {
		return UserName{}, errors.New("user name too long")
	}
	
	return UserName{value: trimmed}, nil
}

func MustUserName(value string) UserName {
	name, err := NewUserName(value)
	if err != nil {
		panic(err)
	}
	return name
}

func (u UserName) Value() string {
	return u.value
}

func (u UserName) String() string {
	return u.value
}

func (u UserName) IsEmpty() bool {
	return u.value == ""
}