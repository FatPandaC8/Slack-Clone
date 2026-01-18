package http

import (
	"auth/internal/auth"
	"auth/internal/database"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users *db.UserRepository
	tokens *db.RefreshTokenRepo
	jwt *auth.JWTService
}

func NewAuthService(u *db.UserRepository, t *db.RefreshTokenRepo, j *auth.JWTService) *AuthService {
	return &AuthService{
		users: u,
		tokens: t,
		jwt: j,
	}
}

func (s *AuthService) Register(name, email, password string) (string, string, string, error) {
	if _, err := s.users.FindByEmail(email); err == nil {
		return "", "", "", errors.New("email already exists")
	}
	hash, _ := auth.HashPassword(password)
	userID := uuid.NewString()
	err := s.users.Create(userID, name, email, hash)
	if err != nil {
		return "", "", "", err
	}

	access, _ := s.jwt.CreateAccessToken(userID)
	refresh, _ := s.tokens.Save(userID)

	return userID, access, refresh, nil
}

func (s *AuthService) Login(
	email, password string,
) (string, string, string, error) {

	u, err := s.users.FindByEmail(email)
	if err != nil {
		return "", "", "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password()), // stored hash
		[]byte(password),     // raw password
	); err != nil {
		return "", "", "", errors.New("invalid credentials")
	}

	access, _ := s.jwt.CreateAccessToken(u.ID())
	refresh, _ := s.tokens.Save(u.ID())

	return u.ID(), access, refresh, nil
}
