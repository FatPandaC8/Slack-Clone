package auth

import (
	"chat-core-go/domain/identity"
	"chat-core-go/domain/valueobject"
	"chat-core-go/ports/service"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthenticator implements the Authenticator port using JWT
type JWTAuthenticator struct {
	secret []byte
}

func NewJWTAuthenticator(secret string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret: []byte(secret),
	}
}

// Authenticate extracts Principal from JWT token
func (a *JWTAuthenticator) Authenticate(ctx context.Context, creds service.Credentials) (*identity.Principal, error) {
	// Type assertion - this adapter knows about Bearer tokens
	bearer, ok := creds.(service.BearerCredentials)
	if !ok {
		return nil, errors.New("invalid credentials type: expected bearer token")
	}
	
	// Parse and validate JWT
	token, err := jwt.Parse(bearer.Token, func(t *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return a.secret, nil
	})
	
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	
	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	
	// Extract user ID
	userIDStr, ok := claims["userId"].(string)
	if !ok {
		return nil, errors.New("missing userId in token")
	}
	userID, err := valueobject.NewUserID(userIDStr)
	if err != nil {
		return nil, errors.New("invalid userId in token")
	}
	
	// Extract expiry
	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("missing expiry in token")
	}
	expiresAt := time.Unix(int64(expFloat), 0)
	
	// Extract roles (optional - default to USER)
	var roles []identity.Role
	if rolesClaim, ok := claims["roles"].([]interface{}); ok {
		for _, r := range rolesClaim {
			if roleStr, ok := r.(string); ok {
				roles = append(roles, identity.Role(roleStr))
			}
		}
	}
	if len(roles) == 0 {
		roles = []identity.Role{identity.RoleUser}
	}
	
	// Create Principal (domain object)
	return identity.NewPrincipal(userID, roles, "jwt-service", expiresAt)
}

// IssueToken generates a JWT from a Principal
func (a *JWTAuthenticator) IssueToken(principal *identity.Principal) (string, error) {
	if principal == nil {
		return "", errors.New("principal cannot be nil")
	}
	
	// Create claims
	claims := jwt.MapClaims{
		"userId": principal.UserID().Value(),
		"roles":  principal.Roles(),
		"iss":    "jwt-service",
		"iat":    principal.IssuedAt().Unix(),
		"exp":    principal.ExpiresAt().Unix(),
	}
	
	// Create and sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}