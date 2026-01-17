package out

import "core/domain/user"

type UserRepository interface {
	FindByID(userID string) (*user.User, error)
}