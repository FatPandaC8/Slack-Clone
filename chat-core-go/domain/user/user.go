package user

import (
	"chat-core-go/domain/valueobject"
	"errors"
	"time"
)

// User is an ENTITY (has identity and lifecycle)
type User struct {
	id           valueobject.UserID
	name         valueobject.UserName
	email        valueobject.Email
	passwordHash string // Not a value object - it's already hashed
	createdAt    time.Time
}

// NewUser creates a new user (factory method)
func NewUser(
	id valueobject.UserID,
	name valueobject.UserName,
	email valueobject.Email,
	passwordHash string,
) (*User, error) {
	if id.IsEmpty() {
		return nil, errors.New("user ID is required")
	}
	if name.IsEmpty() {
		return nil, errors.New("user name is required")
	}
	if email.IsEmpty() {
		return nil, errors.New("email is required")
	}
	if passwordHash == "" {
		return nil, errors.New("password hash is required")
	}
	
	return &User{
		id:           id,
		name:         name,
		email:        email,
		passwordHash: passwordHash,
		createdAt:    time.Now(),
	}, nil
}

// ReconstructUser recreates a user from persistence (no validation)
// Used by repository adapters when loading from database
func ReconstructUser(
	id valueobject.UserID,
	name valueobject.UserName,
	email valueobject.Email,
	passwordHash string,
	createdAt time.Time,
) *User {
	return &User{
		id:           id,
		name:         name,
		email:        email,
		passwordHash: passwordHash,
		createdAt:    createdAt,
	}
}

// Getters
func (u *User) ID() valueobject.UserID {
	return u.id
}

func (u *User) Name() valueobject.UserName {
	return u.name
}

func (u *User) Email() valueobject.Email {
	return u.email
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) ChangeName(newName valueobject.UserName) error {
	if newName.IsEmpty() {
		return errors.New("name cannot be empty")
	}
	u.name = newName
	return nil
}

func (u *User) ChangePassword(newPasswordHash string) error {
	if newPasswordHash == "" {
		return errors.New("password hash cannot be empty")
	}
	u.passwordHash = newPasswordHash
	return nil
}

// Equals checks if two users are the same (based on ID)
func (u *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return u.id.Equals(other.id)
}