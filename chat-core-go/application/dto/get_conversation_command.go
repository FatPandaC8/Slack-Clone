package dto

type ConversationDTO struct {
	ID       string
	Members  []UserDTO
	Messages []MessageDTO
}

type MessageDTO struct {
	ID        string
	SenderID  string
	Content   string
	CreatedAt string
}

type UserDTO struct {
	ID   string
	Name string
}
