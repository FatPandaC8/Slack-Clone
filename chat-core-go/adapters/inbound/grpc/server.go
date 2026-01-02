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
	createChannel in.CreateChannelPort
}

func NewServer(
	sendMessage in.SendMessagePort,
	createChannel in.CreateChannelPort,
) *Server {
	return &Server{
		sendMessage:   sendMessage,
		createChannel: createChannel,
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

func (s *Server) CreateChannel(
	ctx context.Context,
	req *chatpb.CreateChannelRequest,
) (*chatpb.CreateChannelResponse, error) {
	
	var members []user.ID
	for _, id := range req.GetMemberIds() {
		members = append(members, user.ID(id))
	}

	err := s.createChannel.Execute(dto.CreateChannelCommand{
		ChannelID: req.GetChannelId(),
		Members:   members,
	})

	if err != nil {
		return &chatpb.CreateChannelResponse{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	return &chatpb.CreateChannelResponse{Ok: true}, nil
}