// main.go
package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	chatpb "chat-core-go/adapters/inbound/chat-proto"
	grpcadapter "chat-core-go/adapters/inbound/grpc"
	"chat-core-go/config"
)

func main() {
	// Initialize dependency container
	container := config.NewContainer()
	
	// Create gRPC server with auth interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcadapter.AuthInterceptor(container.Authenticator),
		),
	)
	
	// Create gRPC service implementation
	chatServer := grpcadapter.NewServer(
		container.SendMessage,
		container.CreateConversation,
		container.GetConversation,
		container.JoinConversation,
		container.RegisterUser,
		container.LoginUser,
	)
	
	// Register service
	chatpb.RegisterChatServiceServer(grpcServer, chatServer)
	
	// Start listening
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	log.Println(" Chat Core Service (Go) running on :50051")
	
	// Start server (blocking)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}