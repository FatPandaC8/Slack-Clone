package dto

type GetConversationDTO struct {
	ID       string
	Members  []UserDTO
	Messages []MessageDTO
}

type MessageDTO struct {
    ID         string `json:"id"`
    SenderID   string `json:"senderId"`
    SenderName string `json:"senderName,omitempty"`
    Text       string `json:"text"`
    CreatedAt  string `json:"createdAt"`
}

type UserDTO struct {
	ID   string
	Name string
}
