package grpcadapter

import (
	"context"
	"time"

	chatpb "chat-core-go/adapters/inbound/chat-proto"
	"chat-core-go/application/dto"
	"chat-core-go/application/usecase"
	"chat-core-go/domain/valueobject"
)

type Server struct {
	chatpb.UnimplementedChatServiceServer
	
	// Use cases (application layer)
	sendMessage        *usecase.SendMessage
	createConversation *usecase.CreateConversation
	getConversation    *usecase.GetConversation
	joinConversation   *usecase.JoinConversation
}

func NewServer(
	sendMessage *usecase.SendMessage,
	createConversation *usecase.CreateConversation,
	getConversation *usecase.GetConversation,
	joinConversation *usecase.JoinConversation,
) *Server {
	return &Server{
		sendMessage:        sendMessage,
		createConversation: createConversation,
		getConversation:    getConversation,
		joinConversation:   joinConversation,
	}
}

// SendMessage handles sending a message
func (s *Server) SendMessage(
	ctx context.Context,
	req *chatpb.SendMessageRequest,
) (*chatpb.SendMessageResponse, error) {
	
	// Extract Principal from context (injected by interceptor)
	principal, err := GetPrincipal(ctx)
	if err != nil {
		return &chatpb.SendMessageResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}
	
	// Parse value objects from request
	conversationID, err := valueobject.NewConversationID(req.GetConversationId())
	if err != nil {
		return &chatpb.SendMessageResponse{
			Ok:    false,
			Error: "invalid conversation ID: " + err.Error(),
		}, nil
	}
	
	content, err := valueobject.NewMessageContent(req.GetText())
	if err != nil {
		return &chatpb.SendMessageResponse{
			Ok:    false,
			Error: "invalid message content: " + err.Error(),
		}, nil
	}
	
	// Create command
	cmd := dto.SendMessageCommand{
		Principal:      principal,
		ConversationID: conversationID,
		Content:        content,
	}
	
	// Execute use case
	err = s.sendMessage.Execute(cmd)
	if err != nil {
		return &chatpb.SendMessageResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}
	
	return &chatpb.SendMessageResponse{
		Ok:        true,
		CreatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

// CreateConversation handles creating a new conversation
func (s *Server) CreateConversation(
	ctx context.Context,
	req *chatpb.CreateConversationRequest,
) (*chatpb.CreateConversationResponse, error) {
	
	principal, err := GetPrincipal(ctx)
	if err != nil {
		return &chatpb.CreateConversationResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}
	
	name, err := valueobject.NewConversationName(req.GetName())
	if err != nil {
		return &chatpb.CreateConversationResponse{
			Ok:    false,
			Error: "invalid name: " + err.Error(),
		}, nil
	}
	
	cmd := dto.CreateConversationCommand{
		Principal: principal,
		Name:      name,
	}
	
	result, err := s.createConversation.Execute(cmd)
	if err != nil {
		return &chatpb.CreateConversationResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}
	
	return &chatpb.CreateConversationResponse{
		Ok:             true,
		ConversationId: result.ConversationID.Value(),
		Name:           result.Name.Value(),
		InviteCode:     result.InviteCode.Value(),
	}, nil
}

// GetConversation retrieves conversation details
func (s *Server) GetConversation(
	ctx context.Context,
	req *chatpb.GetConversationRequest,
) (*chatpb.GetConversationResponse, error) {
	
	principal, err := GetPrincipal(ctx)
	if err != nil {
		return nil, err
	}
	
	conversationID, err := valueobject.NewConversationID(req.GetConversationId())
	if err != nil {
		return nil, err
	}
	
	query := dto.GetConversationQuery{
		Principal:      principal,
		ConversationID: conversationID,
	}
	
	result, err := s.getConversation.Execute(query)
	if err != nil {
		return nil, err
	}
	
	// Convert to protobuf response
	response := &chatpb.GetConversationResponse{
		ConversationId: result.ConversationID.Value(),
	}
	
	// Add messages
	for _, msg := range result.Messages {
		response.Messages = append(response.Messages, &chatpb.ChatMessage{
			MessageId: msg.MessageID.Value(),
			SenderId:  msg.SenderID.Value(),
			Name:      msg.SenderName.Value(),
			Text:      msg.Content.Value(),
			CreatedAt: msg.CreatedAt.Format(time.RFC3339),
		})
	}
	
	// Add members
	for _, member := range result.Members {
		response.Members = append(response.Members, &chatpb.ChatUser{
			UserId: member.UserID.Value(),
			Name:   member.Name.Value(),
		})
	}
	
	return response, nil
}

// JoinConversation handles joining via invite code
func (s *Server) JoinConversation(
	ctx context.Context,
	req *chatpb.JoinConversationRequest,
) (*chatpb.JoinConversationResponse, error) {
	
	principal, err := GetPrincipal(ctx)
	if err != nil {
		return &chatpb.JoinConversationResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}
	
	inviteCode, err := valueobject.NewInviteCode(req.GetInviteCode())
	if err != nil {
		return &chatpb.JoinConversationResponse{
			Ok:    false,
			Error: "invalid invite code: " + err.Error(),
		}, nil
	}
	
	cmd := dto.JoinConversationCommand{
		Principal:  principal,
		InviteCode: inviteCode,
	}
	
	err = s.joinConversation.Execute(cmd)
	if err != nil {
		return &chatpb.JoinConversationResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}
	
	return &chatpb.JoinConversationResponse{
		Ok: true,
	}, nil
}