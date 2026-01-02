package in

import "chat-core-go/domain/conversation"

type GetConversationPort interface {
	Execute(conversationID string) (*conversation.Conversation, error)
}