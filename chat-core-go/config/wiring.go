package config

import (
	"chat-core-go/adapters/outbound/bcryptadapter"
	"chat-core-go/adapters/outbound/jwtadapter"
	"chat-core-go/adapters/outbound/persistent/inmemory"
	"chat-core-go/adapters/outbound/publisher"
	"chat-core-go/application/usecase"
)

var conversationRepo = persistent.NewInMemoryConversationRepo()
var messageRepo = persistent.NewInMemoryMessageRepo()
var userRepo = persistent.NewInmemoryUserRepo()
var passwordHasher = bcryptadapter.New()
var TokenJWT = jwtadapter.NewJWTService("super-secret-dev-key") // for production, put it in env variables

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

func WireListUsers() *usecase.ListUsers {
	return usecase.NewListUsers(userRepo)
}

func WireJoinConversation() *usecase.JoinConversation {
	return usecase.NewJoinConversation(conversationRepo)
}

func WireRegisterUser() *usecase.RegisterUser {
	return usecase.NewRegisterUser(
		userRepo,
		passwordHasher,
	)
}

func WireLoginUser() *usecase.LoginUser {
	return usecase.NewLoginUser(
		userRepo,
		passwordHasher,
		TokenJWT,
	)
}