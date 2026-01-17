package message

import "time"

type Message struct {
	messageID 		string
	roomID    string
	senderID string
	content   string
	createdAt time.Time
}


func NewMessage(id, roomID, senderID, content string, createdAt time.Time) *Message {
	return &Message{
		messageID:        id,
		roomID:    roomID,
		senderID: senderID,
		content:   content,
		createdAt: createdAt,
	}
}

func (m *Message) ID() string {
	return m.messageID
}

func (m *Message) RoomID() string {
	return m.roomID
}

func (m *Message) SenderID() string {
	return m.senderID
}

func (m *Message) Content() string {
	return m.content
}

func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}