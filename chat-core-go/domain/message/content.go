package message

type Content struct {
	value string
}

func NewContent(text string) Content {
	if text == "" {
		panic("message cannot be empty")
	} 
	if len(text) > 500 {
		panic("message too long")
	}
	return Content{value: text}
}

// value getter
func (c Content) Value() string {
	return c.value
}