// domain/identity/principal.go
package identity

import (
	"chat-core-go/domain/valueobject"
	"errors"
	"time"
)

// Principal represents an authenticated entity in the system.
// It's a VALUE OBJECT - immutable and transport-agnostic.
type Principal struct {
	userID    valueobject.UserID
	roles     []Role
	issuer    string
	issuedAt  time.Time
	expiresAt time.Time
}

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

// NewPrincipal creates a validated Principal
func NewPrincipal(
	userID valueobject.UserID,
	roles []Role,
	issuer string,
	issuedAt time.Time,
	expiresAt time.Time,
) (*Principal, error) {
	if userID.IsEmpty() {
		return nil, errors.New("user ID cannot be empty")
	}
	if issuer == "" {
		return nil, errors.New("issuer cannot be empty")
	}
	if len(roles) == 0 {
		roles = []Role{RoleUser} // default role
	}
	
	return &Principal{
		userID:    userID,
		roles:     roles,
		issuer:    issuer,
		issuedAt:  issuedAt,
		expiresAt: expiresAt,
	}, nil
}

// SystemPrincipal creates a principal for system-generated actions
func SystemPrincipal() *Principal {
	return &Principal{
		userID:    valueobject.MustUserID("system"),
		roles:     []Role{RoleAdmin},
		issuer:    "system",
		issuedAt:  time.Now(),
		expiresAt: time.Now().Add(24 * time.Hour),
	}
}

// Getters (immutable - no setters!)
func (p *Principal) UserID() valueobject.UserID {
	return p.userID
}

func (p *Principal) Roles() []Role {
	// Return copy to prevent modification
	rolesCopy := make([]Role, len(p.roles))
	copy(rolesCopy, p.roles)
	return rolesCopy
}

func (p *Principal) Issuer() string {
	return p.issuer
}

func (p *Principal) IssuedAt() time.Time {
	return p.issuedAt
}

func (p *Principal) ExpiresAt() time.Time {
	return p.expiresAt
}

// Domain rules encoded in Principal
func (p *Principal) HasRole(role Role) bool {
	for _, r := range p.roles {
		if r == role {
			return true
		}
	}
	return false
}

func (p *Principal) IsExpired() bool {
	return time.Now().After(p.expiresAt)
}

func (p *Principal) IsAdmin() bool {
	return p.HasRole(RoleAdmin)
}

func (p *Principal) IsUser() bool {
	return p.HasRole(RoleUser)
}

func (p *Principal) IsSystem() bool {
	return p.issuer == "system"
}