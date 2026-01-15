package out

import (
	"chat-core-go/domain/user"
	"chat-core-go/domain/valueobject"
)

type UserRepository interface {
	// Save persists a user
	Save(user *user.User) error
	
	// Load retrieves a user by ID
	Load(id valueobject.UserID) (*user.User, error)
	
	// LoadByEmail retrieves a user by email
	LoadByEmail(email valueobject.Email) (*user.User, error)
	
	// List returns all users (for admin purposes)
	List() ([]*user.User, error)
	
	// GenerateID creates a new unique user ID
	GenerateID() valueobject.UserID
	
	// Exists checks if a user exists
	Exists(id valueobject.UserID) (bool, error)
}