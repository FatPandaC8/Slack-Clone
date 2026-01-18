package db

import (
	"auth/internal/user"
	"database/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB	
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// take email, password and return user id
func (u *UserRepository) Create(userid, name, email, password string) error {
	err := u.db.QueryRow(
		`INSERT INTO users (userid, name, email, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING userid`,
		userid, name, email, password, time.Now(),
	).Scan(&userid)

	return err
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var (
		userID    string
		name      string
		password  string
		createdAt time.Time
	)

	err := r.db.QueryRow(
		`SELECT userid, name, password_hash, created_at
		 FROM users
		 WHERE email = $1`,
		email,
	).Scan(&userID, &name, &password, &createdAt)

	if err != nil {
		return nil, err
	}

	return user.NewUser(
		userID,
		name,
		email,
		password,
		createdAt,
	), nil
}
