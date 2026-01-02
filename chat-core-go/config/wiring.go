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

func WireCreateChannel() *usecase.CreateChannel {
	return usecase.NewCreateChannel(conversationRepo)
}
