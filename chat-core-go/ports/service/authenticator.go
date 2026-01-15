package service

import (
	"chat-core-go/domain/identity"
	"context"
)

// Authenticator extracts and validates identity from credentials.
// This is the PORT - the domain defines WHAT it needs, not HOW.
type Authenticator interface {
	// Authenticate validates credentials and returns a Principal
	Authenticate(ctx context.Context, credentials Credentials) (*identity.Principal, error)
}

// Credentials is an opaque interface - adapters provide implementations
type Credentials interface {
	Type() string
}

// TokenIssuer generates tokens for authenticated principals
type TokenIssuer interface {
	// IssueToken creates a token for a principal
	IssueToken(principal *identity.Principal) (string, error)
}

// ----- Credential Implementations (used by adapters) -----

// BearerCredentials represents a Bearer token (JWT, OAuth, etc.)
type BearerCredentials struct {
	Token string
}

func (b BearerCredentials) Type() string {
	return "bearer"
}