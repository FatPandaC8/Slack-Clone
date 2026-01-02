package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	chatpb "chat-core-go/adapters/inbound/chat-proto"
	grpcadapter "chat-core-go/adapters/inbound/grpc"
	"chat-core-go/config"
	"fmt"
)

func main() {
	sendMessageUC := config.WireSendMessage()
	createMessageUC := config.WireCreateChannel()

	grpcServer := grpc.NewServer()

	chatServer := grpcadapter.NewServer(sendMessageUC, createMessageUC)
	chatpb.RegisterChatServiceServer(grpcServer, chatServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 chat-core-go gRPC running on :50051")

	// 5. Start server (blocking)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

	fmt.Println("[CORE] Running...")
}
