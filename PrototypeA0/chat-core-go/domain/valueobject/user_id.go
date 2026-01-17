// domain/valueobject/user_id.go
package valueobject

import "errors"

// UserID is a value object representing a user identifier.
// It's NOT just a string - it carries domain rules.
type UserID struct {
	value string
}

// NewUserID creates a validated UserID
func NewUserID(value string) (UserID, error) {
	if value == "" {
		return UserID{}, errors.New("user ID cannot be empty")
	}
	if len(value) > 100 {
		return UserID{}, errors.New("user ID too long")
	}
	return UserID{value: value}, nil
}

// MustUserID creates a UserID or panics (for tests/constants)
func MustUserID(value string) UserID {
	id, err := NewUserID(value)
	if err != nil {
		panic(err)
	}
	return id
}

// Value returns the underlying string (read-only)
func (u UserID) Value() string {
	return u.value
}

// Equals checks equality (value object comparison)
func (u UserID) Equals(other UserID) bool {
	return u.value == other.value
}

// String implements Stringer
func (u UserID) String() string {
	return u.value
}

// IsEmpty checks if this is a zero-value UserID
func (u UserID) IsEmpty() bool {
	return u.value == ""
}