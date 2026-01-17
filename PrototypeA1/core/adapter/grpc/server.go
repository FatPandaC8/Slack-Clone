package grpcadapter

import (
	"context"
	pb "core/adapter/grpc/proto"
	"core/application"
	"core/port/in"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)


type Server struct {
	pb.UnimplementedChatServiceServer

	createRoomUC application.CreateRoomUseCase
	joinRoomUC application.JoinRoomUseCase
	sendMessageUC application.SendMessageUseCase
	listMessagesUC application.ListMessagesUseCase
}

func NewServer(
	createRoomUC application.CreateRoomUseCase,
	joinRoomUC application.JoinRoomUseCase,
	sendMessageUC application.SendMessageUseCase,
	listMessagesUC application.ListMessagesUseCase,
) *Server {
	return &Server{
		createRoomUC:   createRoomUC,
		joinRoomUC:     joinRoomUC,
		sendMessageUC:  sendMessageUC,
		listMessagesUC: listMessagesUC,
	}
}

func (s *Server) CreateRoom(
	ctx context.Context,
	req *pb.CreateRoomRequest,
) (*pb.CreateRoomResponse, error) {
	res, err := s.createRoomUC.CreateRoom(in.CreateRoomCommand{
		UserID: req.UserID,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomResponse{
		RoomID: res.RoomID,
		InviteCode: res.InviteCode,
	}, nil
}

func (s *Server) JoinRoom(
	ctx context.Context,
	req *pb.JoinRoomRequest,
) (*pb.JoinRoomResponse, error) {

	err := s.joinRoomUC.JoinRoom(in.JoinRoomCommand{
		UserID:     req.UserID,
		InviteCode: req.InviteCode,
	})
	if err != nil {
		return nil, err
	}

	return &pb.JoinRoomResponse{}, nil
}

func (s *Server) SendMessage(
	ctx context.Context,
	req *pb.SendMessageRequest,
) (*pb.SendMessageResponse, error) {

	err := s.sendMessageUC.SendMessage(in.SendMessageCommand{
		UserID:  req.UserID,
		RoomID:  req.RoomID,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{}, nil
}

func (s *Server) ListMessages(
	ctx context.Context,
	req *pb.ListMessagesRequest,
) (*pb.ListMessagesResponse, error) {

	messages, err := s.listMessagesUC.ListMessages(in.ListMessagesQuery{
		UserID: req.UserID,
		RoomID: req.RoomID,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*pb.Message, 0, len(messages))
	for _, m := range messages {
		result = append(result, &pb.Message{
			MessageID:        m.ID,
			SenderID: m.SenderID,
			Content:  m.Content,
			CreatedAt: timestamppb.New(
				time.Time(m.Created),
			),
		})
	}

	return &pb.ListMessagesResponse{
		Messages: result,
	}, nil
}
