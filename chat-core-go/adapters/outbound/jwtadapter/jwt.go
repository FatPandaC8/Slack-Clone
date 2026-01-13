package jwtadapter

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secret: []byte(secret),
	}
}

func (s *JWTService) GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *JWTService) VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func (t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	id, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("no user id")
	}
	return id, nil
}