package postgres

import (
	"core/domain/user"
	valueobject "core/domain/valueobject/user"
	"database/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(userID valueobject.UserID) (*user.User, error) {
	row := r.db.QueryRow(
		`SELECT userid, name, email, createdat
		 FROM users
		 WHERE id = $1`,
		userID.String(),
	)

	var (
		rawUserID    string
		name      string
		email     string
		createdAt time.Time
	)

	if err := row.Scan(&rawUserID, &name, &email, &createdAt); err != nil {
		return nil, err
	}

	return user.NewUser(userID, name, email, createdAt), nil
}
