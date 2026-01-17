package postgres

import (
	"core/domain/user"
	"database/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id string) (*user.User, error) {
	row := r.db.QueryRow(
		`SELECT id, name, email, created_at
		 FROM users
		 WHERE id = $1`,
		id,
	)

	var (
		userID    string
		name      string
		email     string
		createdAt time.Time
	)

	if err := row.Scan(&userID, &name, &email, &createdAt); err != nil {
		return nil, err
	}

	return user.NewUser(userID, name, email, createdAt), nil
}
