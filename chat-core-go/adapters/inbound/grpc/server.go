package grpcadapter

import (
	"context"

	chatpb "chat-core-go/adapters/inbound/chat-proto"
	"chat-core-go/application/dto"
	"chat-core-go/domain/user"
	"chat-core-go/ports/in"
)

type Server struct {
	chatpb.UnimplementedChatServiceServer
	sendMessage in.SendMessagePort
	createConversation in.CreateConversationPort
	getConversation in.GetConversationPort
	listConversations in.ListConversationPort
}

func NewServer(
	sendMessage in.SendMessagePort,
	createConversation in.CreateConversationPort,
	getConversation in.GetConversationPort,
	listConversations in.ListConversationPort,
) *Server {
	return &Server{
		sendMessage:   sendMessage,
		createConversation: createConversation,
		getConversation: getConversation,
		listConversations: listConversations,
	}
}


func (s *Server) SendMessage(
	ctx context.Context,
	req *chatpb.SendMessageRequest,
) (*chatpb.SendMessageResponse, error) {

	cmd := dto.SendMessageCommand{
		MessageID: 		req.GetMessageId(),
		ConversationID: req.GetConversationId(),
		SenderID:       user.ID(req.GetSenderId()),
		Text:        	req.GetText(),
	}

	err := s.sendMessage.Execute(cmd)
	if err != nil {
		return &chatpb.SendMessageResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	return &chatpb.SendMessageResponse{
		Ok: true,
	}, nil
}

func (s *Server) CreateConversation(
	ctx context.Context,
	req *chatpb.CreateConversationRequest,
) (*chatpb.CreateConversationResponse, error) {
	
	var members []user.ID
	for _, id := range req.GetMemberIds() {
		members = append(members, user.ID(id))
	}

	err := s.createConversation.Execute(dto.CreateConversationCommand{
		ConversationID: req.GetConversationId(),
		Members:   members,
	})

	if err != nil {
		return &chatpb.CreateConversationResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	return &chatpb.CreateConversationResponse{Ok: true}, nil
}

func (s *Server) GetConversation(
	ctx context.Context,
	req *chatpb.GetConversationRequest,
) (*chatpb.GetConversationResponse, error) {
	conv, err := s.getConversation.Execute(req.GetConversationId())
	if err != nil {
		return nil, err
	}

	res := &chatpb.GetConversationResponse{
		ConversationId: string(conv.ID()),
	}

	for _, msg := range conv.Messages() {
		res.Messages = append(res.Messages, &chatpb.ChatMessage{
			MessageId: string(msg.ID()),
			SenderId: string(msg.Sender()),
			Text: msg.Content().Value(),
		})
	}

	return res, nil
}

func (s *Server) ListConversations(
	ctx context.Context, 
	req *chatpb.ListConversationsRequest,
) (*chatpb.ListConversationsResponse, error) {
	convs, err := s.listConversations.Execute(user.ID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	res := &chatpb.ListConversationsResponse{}
	for _, c := range convs {
		res.ConversationId = append(
			res.ConversationId, 
			string(c.ID()),
		)
	}

	return res, nil
}