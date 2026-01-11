package grpcadapter

import (
	"context"

	chatpb "chat-core-go/adapters/inbound/chat-proto"
	"chat-core-go/application/dto"
	"chat-core-go/ports/in"
)

type Server struct {
	chatpb.UnimplementedChatServiceServer
	sendMessage        in.SendMessagePort
	createConversation in.CreateConversationPort
	getConversation    in.GetConversationPort
	listConversations  in.ListConversationPort
	createUser         in.CreateUserPort
	listUser           in.ListUserPort
	joinConversation   in.JoinConversationPort
}

func NewServer(
	sendMessage in.SendMessagePort,
	createConversation in.CreateConversationPort,
	getConversation in.GetConversationPort,
	listConversations in.ListConversationPort,
	createUser in.CreateUserPort,
	listUser in.ListUserPort,
	joinConversation in.JoinConversationPort,
) *Server {
	return &Server{
		sendMessage:        sendMessage,
		createConversation: createConversation,
		getConversation:    getConversation,
		listConversations:  listConversations,
		createUser:         createUser,
		listUser:           listUser,
		joinConversation: joinConversation,
	}
}

func (s *Server) CreateUser(
	ctx context.Context,
	req *chatpb.CreateUserRequest,
) (*chatpb.CreateUserResponse, error) {

	cmd := dto.CreateUserCommand{
		Name:         req.GetName(),
		Email:        req.GetEmail(),
		PasswordHash: req.GetPassword(),
	}

	u, err := s.createUser.Execute(cmd)
	if err != nil {
		return &chatpb.CreateUserResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	return &chatpb.CreateUserResponse{
		Ok:     true,
		UserId: u.ID,
		Name:   u.Name,
	}, nil
}

func (s *Server) ListUsers(
	ctx context.Context,
	req *chatpb.ListUserRequest,
) (*chatpb.ListUserResponse, error) {

	users, err := s.listUser.Execute(req.GetConversationId())
	if err != nil {
		return nil, err
	}

	res := &chatpb.ListUserResponse{}
	for _, u := range users {
		res.Users = append(res.Users, &chatpb.ChatUser{
			UserId: u.ID,
			Name:   u.Name,
		})
	}

	return res, nil
}

func (s *Server) SendMessage(
	ctx context.Context,
	req *chatpb.SendMessageRequest,
) (*chatpb.SendMessageResponse, error) {

	cmd := dto.SendMessageCommand{
		MessageID:      req.GetMessageId(),
		ConversationID: req.GetConversationId(),
		SenderID:       req.GetSenderId(),
		Text:           req.GetText(),
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

	result, err := s.createConversation.Execute(dto.CreateConversationCommand{
		Name: req.GetName(),
		CreatorID: req.GetCreatorId(),
	})
	if err != nil {
		return &chatpb.CreateConversationResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	res := &chatpb.CreateConversationResponse{
		Ok:             true,
		ConversationId: result.ID,
		InviteCode: result.InviteCode,
		Name: result.Name,
	}

	return res, nil
}

func (s *Server) JoinConversation(ctx context.Context, req *chatpb.JoinConversationRequest) (
    *chatpb.JoinConversationResponse, error) {

    err := s.joinConversation.Execute(dto.JoinConversationCommand{
        InviteCode: req.GetInviteCode(),
        UserID: req.GetUserId(),
    })

    if err != nil {
        return &chatpb.JoinConversationResponse{
            Ok: false,
            Error: err.Error(),
        }, nil
    }

    return &chatpb.JoinConversationResponse{
        Ok: true,
    }, nil
}


func (s *Server) GetConversation(
	ctx context.Context,
	req *chatpb.GetConversationRequest,
) (*chatpb.GetConversationResponse, error) {

	convDTO, err := s.getConversation.Execute(req.GetConversationId())
	if err != nil {
		return nil, err
	}

	res := &chatpb.GetConversationResponse{
		ConversationId: convDTO.ID,
	}

	for _, msg := range convDTO.Messages {
		res.Messages = append(res.Messages, &chatpb.ChatMessage{
			MessageId: msg.ID,
			SenderId:  msg.SenderID,
			Text:      msg.Content,
			Name: 	msg.Name,	
			CreatedAt: msg.CreatedAt,
		})
	}

	for _, member := range convDTO.Members {
		res.Members = append(res.Members, &chatpb.ChatUser{
			UserId: member.ID,
			Name:   member.Name,
		})
	}

	return res, nil
}

func (s *Server) ListConversations(
	ctx context.Context,
	req *chatpb.ListConversationsRequest,
) (*chatpb.ListConversationsResponse, error) {
	convs, err := s.listConversations.Execute(req.GetUserId())
	if err != nil {
		return nil, err
	}

	res := &chatpb.ListConversationsResponse{}
	for _, c := range convs {
		res.Conversations = append(
			res.Conversations,
			&chatpb.ConversationView{
				ConversationId: c.ID(),
				Name: c.Name(),
			},
		)
	}

	return res, nil
}

