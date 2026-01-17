package main

import (
	grpcadapter "core/adapter/grpc"
	pb "core/adapter/grpc/proto"
	"core/config"
	"log"
	"net"

	"google.golang.org/grpc"
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

	grpcServer := grpc.NewServer()
	chatServer := grpcadapter.NewServer(
		app.CreateRoom,
		app.JoinRoom,
		app.SendMessage,
		app.ListMessage,
	)
	pb.RegisterChatServiceServer(grpcServer, chatServer)
	
	log.Println("gRPC server listening on", cfg.GRPC.Addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}