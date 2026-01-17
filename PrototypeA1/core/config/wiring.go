package config

import (
	"core/adapter/postgres"
	"core/adapter/ws"
	"core/application"
	"database/sql"
)

type App struct {
	DB *sql.DB

	CreateRoom application.CreateRoomUseCase
	JoinRoom application.JoinRoomUseCase
	SendMessage application.SendMessageUseCase
	ListMessage application.ListMessagesUseCase
}

func Wire(cfg Config) (*App, error) {
	db, err := postgres.Open(postgres.Config{
		Host: cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		return nil, err
	}

	roomRepo := postgres.NewRoomRepository(db)
	roomMemberRepo := postgres.NewRoomMemberRepository(db)
	messageRepo := postgres.NewMessageRepository(db)
	broadcaster := ws.NewBroadcaster()

	// --- use cases ---
	createRoom := application.NewCreateRoomUseCase(
		roomRepo,
		roomMemberRepo,
	)

	joinRoom := application.NewJoinRoomUseCase(
		roomRepo,
		roomMemberRepo,
	)

	listMessages := application.NewListMessagesUseCase(
		roomMemberRepo,
		messageRepo,
	)

	sendMessage := application.NewSendMessageUseCase(
		roomMemberRepo,
		messageRepo,
		broadcaster,
	)

	return &App{
		DB:           db,
		CreateRoom:   *createRoom,
		JoinRoom:     *joinRoom,
		ListMessage: *listMessages,
		SendMessage:  *sendMessage,
	}, nil
}