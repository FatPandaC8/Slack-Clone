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
	sendMessageUC := config.WireSendMessage()
	createConversationUC := config.WireCreateConversation()
	getConversationUC := config.WireGetConversation()
	listConversationsUC := config.WireListConversations()

	grpcServer := grpc.NewServer()

	chatServer := grpcadapter.NewServer(
		sendMessageUC,
		createConversationUC,
		getConversationUC, 
		listConversationsUC,
	)
	
	chatpb.RegisterChatServiceServer(grpcServer, chatServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("chat-core-go gRPC running on :50051")

	// 5. Start server (blocking)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
