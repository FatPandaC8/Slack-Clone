package out

import (
	"core/domain/user"
	valueobject "core/domain/valueobject/user"
)

type UserRepository interface {
	FindByID(userID valueobject.UserID) (*user.User, error)
}