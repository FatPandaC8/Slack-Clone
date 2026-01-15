// domain/valueobject/email.go
package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

// Email represents a validated email address
type Email struct {
	value string
}

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(value string) (Email, error) {
	normalized := strings.TrimSpace(strings.ToLower(value))
	
	if normalized == "" {
		return Email{}, errors.New("email cannot be empty")
	}
	if !emailPattern.MatchString(normalized) {
		return Email{}, errors.New("invalid email format")
	}
	if len(normalized) > 255 {
		return Email{}, errors.New("email too long")
	}
	
	return Email{value: normalized}, nil
}

func MustEmail(value string) Email {
	email, err := NewEmail(value)
	if err != nil {
		panic(err)
	}
	return email
}

func (e Email) Value() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}

func (e Email) String() string {
	return e.value
}

func (e Email) IsEmpty() bool {
	return e.value == ""
}