package config

import (
	"chat-core-go/adapters/outbound/persistent"
	"chat-core-go/adapters/outbound/publisher"
	"chat-core-go/application/usecase"
)

var conversationRepo = persistent.NewInMemoryConversationRepo()
var messageRepo = persistent.NewInMemoryMessageRepo()

func WireSendMessage() *usecase.SendMessage {
	pub := &publisher.LogPublisher{}
	return usecase.NewSendMessage(conversationRepo, messageRepo, pub)
}

func WireCreateConversation() *usecase.CreateConversation {
	return usecase.NewCreateConversation(conversationRepo)
}

func WireGetConversation() *usecase.GetConversation {
	return usecase.NewGetConversation(conversationRepo)
}