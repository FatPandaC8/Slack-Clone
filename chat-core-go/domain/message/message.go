package message

import (
	"time"
)

type Message struct {
	id 				string
	sender 			string
	conversationID	string
	content 		Content
	createdAt 		time.Time
}

func NewMessage(
	id string,
	sender string,
	conversationID string,
	content Content,
) Message {
	return Message{
		id:             id,
		sender:         sender,
		conversationID: conversationID,
		content:        content,
	}
}

func (m Message) ID() string                { return m.id }
func (m Message) Sender() string       		{ return m.sender }
func (m Message) ConversationID() string 	{ return m.conversationID }
func (m Message) Content() Content      	{ return m.content }
func (m Message) CreatedAt() time.Time 		{ return m.createdAt }