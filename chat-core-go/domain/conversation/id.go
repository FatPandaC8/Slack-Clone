package conversation

type ID string

func NewID(value string) ID {
	if value == "" {
		panic("conversation id cannot be empty")
	}
	return ID(value)
}
