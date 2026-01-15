package config

import (
	"chat-core-go/adapters/outbound/auth"
	"chat-core-go/adapters/outbound/persistent/inmemory"
	"chat-core-go/application/usecase"
	"chat-core-go/ports/out"
	"chat-core-go/ports/service"
)

type Container struct {
	// Repositories
	UserRepo         out.UserRepository
	ConversationRepo out.ConversationRepository
	MessageRepo      out.MessageRepository
	
	// Services
	Authenticator service.Authenticator
	TokenIssuer   service.TokenIssuer
	PasswordHasher service.PasswordHasher
	
	// Use cases
	SendMessage        *usecase.SendMessage
	CreateConversation *usecase.CreateConversation
	GetConversation    *usecase.GetConversation
	JoinConversation   *usecase.JoinConversation
	RegisterUser       *usecase.RegisterUser
	LoginUser          *usecase.LoginUser
}

// NewContainer initializes all dependencies
func NewContainer() *Container {
	// Initialize repositories (adapters)
	userRepo := persistent.NewUserRepository()
	conversationRepo := persistent.NewConversationRepository()
	messageRepo := persistent.NewMessageRepository()
	
	// Initialize services (adapters)
	jwtSecret := "xsaeslrtjtupzegqbzkuohkotelteuxvqlnmwrhonlrvhyvfterihobznadpjttf"
	jwtAuth := auth.NewJWTAuthenticator(jwtSecret)
	bcryptHasher := auth.NewBcryptHasher()
	
	// Initialize use cases (application layer)
	sendMessage := usecase.NewSendMessage(
		conversationRepo,
		messageRepo,
		userRepo,
	)
	
	createConversation := usecase.NewCreateConversation(
		conversationRepo,
	)
	
	getConversation := usecase.NewGetConversation(
		conversationRepo,
		messageRepo,
		userRepo,
	)
	
	joinConversation := usecase.NewJoinConversation(
		conversationRepo,
		userRepo,
	)
	
	registerUser := usecase.NewRegisterUser(
		userRepo,
		bcryptHasher,
	)
	
	loginUser := usecase.NewLoginUser(
		userRepo,
		bcryptHasher,
		jwtAuth, // TokenIssuer interface
	)
	
	return &Container{
		UserRepo:         userRepo,
		ConversationRepo: conversationRepo,
		MessageRepo:      messageRepo,
		Authenticator:    jwtAuth,
		TokenIssuer:      jwtAuth,
		PasswordHasher:   bcryptHasher,
		SendMessage:        sendMessage,
		CreateConversation: createConversation,
		GetConversation:    getConversation,
		JoinConversation:   joinConversation,
		RegisterUser:       registerUser,
		LoginUser:          loginUser,
	}
}