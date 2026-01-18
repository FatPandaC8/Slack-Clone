package db

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"time"
)

type RefreshTokenRepo struct {
	db *sql.DB
}

func NewRefreshTokenRepo(db *sql.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{
		db: db,
	}
}

func generateRefreshToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func (r *RefreshTokenRepo) Save(userID string) (string, error) {
	token := generateRefreshToken()
	exp := time.Now().Add(7 * time.Minute)
	_, err := r.db.Exec(
		`INSERT INTO refresh_tokens (token, userid, expires_at)
		VALUES ($1, $2, $3)`,
		token, userID, exp,
	)
	return token, err
}