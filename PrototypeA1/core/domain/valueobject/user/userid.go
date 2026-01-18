package valueobject

import "errors"

type UserID struct {
	value string
}

func NewUserID(v string) (UserID, error) {
	if v == "" {
		return UserID{}, errors.New("user id is empty")
	}
	return UserID{
		value: v,
	}, nil
}

func (u UserID) String() string {
	return u.value
}