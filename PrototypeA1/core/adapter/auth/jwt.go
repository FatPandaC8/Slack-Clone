package auth

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt"
)

type JWTVerifier struct {
	secret []byte
}

func NewJWTVerifier(secret string) *JWTVerifier {
	return &JWTVerifier{
		secret: []byte(secret),
	}
}

func (j *JWTVerifier) OldVerify(tokenStr string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func (t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return j.secret, nil
	})

	log.Println("TOKEN VALID:", token.Valid)

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	
	return claims["sub"].(string), nil
}

func (j *JWTVerifier) Verify(tokenStr string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secret), nil // ✅ MUST be []byte
	})

	if err != nil {
		log.Println("JWT parse error:", err)
		return "", errors.New("invalid token")
	}

	if !token.Valid {
		log.Println("JWT not valid")
		return "", errors.New("invalid token")
	}

	// SAFE extraction of sub
	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", errors.New("missing sub claim")
	}

	return sub, nil
}
