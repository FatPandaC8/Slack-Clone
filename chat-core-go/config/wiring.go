package config

import (
	"chat-core-go/adapters/outbound/persistent/inmemory"
	"chat-core-go/adapters/outbound/publisher"
	"chat-core-go/application/usecase"
)

var conversationRepo = persistent.NewInMemoryConversationRepo()
var messageRepo = persistent.NewInMemoryMessageRepo()
var userRepo = persistent.NewInmemoryUserRepo()

func WireSendMessage() *usecase.SendMessage {
	pub := &publisher.LogPublisher{}
	return usecase.NewSendMessage(conversationRepo, messageRepo, pub)
}

func WireCreateConversation() *usecase.CreateConversation {
	return usecase.NewCreateConversation(conversationRepo)
}

func WireGetConversation() *usecase.GetConversation {
	return usecase.NewGetConversation(conversationRepo, messageRepo, userRepo)
}

func WireListConversations() *usecase.ListConversations {
	return usecase.NewListConversations(conversationRepo)
}

func WireCreateUser() *usecase.CreateUser {
	return usecase.NewCreateUser(userRepo)
}

func WireListUsers() *usecase.ListUsers {
	return usecase.NewListUsers(userRepo)
}

