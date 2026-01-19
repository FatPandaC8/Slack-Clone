package main

import (
	"core/adapter/auth"
	grpcadapter "core/adapter/grpc"
	pb "core/adapter/grpc/proto"
	"core/config"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	app, err := config.Wire(cfg)
	if err != nil {
		log.Fatal("failed to wire app:", err)
	}
	defer app.DB.Close()

	lis, err := net.Listen("tcp", cfg.GRPC.Addr)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	verfier := auth.NewJWTVerifier(cfg.JWT.Secret)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			auth.UnaryAuthInterceptor(verfier),
		),
	)
	chatServer := grpcadapter.NewServer(
		app.CreateRoom,
		app.JoinRoom,
		app.SendMessage,
		app.ListMessage,
	)
	pb.RegisterChatServiceServer(grpcServer, chatServer)
	reflection.Register(grpcServer)
	
	log.Println("gRPC server listening on", cfg.GRPC.Addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}