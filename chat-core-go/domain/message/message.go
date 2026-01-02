package message

import "chat-core-go/domain/user"

type Message struct {
	id 				ID
	sender 			user.ID
	conversationID	string
	content 		Content
}

func NewMessage(
	id ID,
	sender user.ID,
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

func (m Message) ID() ID                { return m.id }
func (m Message) Sender() user.ID       { return m.sender }
func (m Message) ConversationID() string { return m.conversationID }
func (m Message) Content() Content      { return m.content }