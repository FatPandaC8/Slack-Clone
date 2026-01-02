package message

type ID string

func NewID(value string) ID {
	if value == "" {
		panic("message id cannot be empty")
	}
	return ID(value)
}