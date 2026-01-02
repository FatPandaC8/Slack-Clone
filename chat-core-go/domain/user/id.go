package user

type ID string

func NewID(value string) ID {
	if value == "" {
		panic("user id cannot be empty")
	}
	return ID(value)
}