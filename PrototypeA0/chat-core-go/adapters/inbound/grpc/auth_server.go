package grpcadapter

import (
	"context"
	"time"

	authpb "chat-core-go/adapters/inbound/chat-proto"
	"chat-core-go/application/dto"
	"chat-core-go/application/usecase"
	"chat-core-go/domain/identity"
	"chat-core-go/domain/valueobject"
	"chat-core-go/ports/service"
)

type AuthServer struct {
	authpb.UnimplementedChatServiceServer

	registerUser *usecase.RegisterUser
	loginUser    *usecase.LoginUser
	tokenIssuer  service.TokenIssuer
}

func NewAuthServer(
	registerUser *usecase.RegisterUser,
	loginUser *usecase.LoginUser,
	tokenIssuer service.TokenIssuer,
) *AuthServer {
	return &AuthServer{
		registerUser: registerUser,
		loginUser:    loginUser,
		tokenIssuer:  tokenIssuer,
	}
}

func (s *AuthServer) RegisterUser(
	ctx context.Context,
	req *authpb.RegisterUserRequest,
) (*authpb.RegisterUserResponse, error) {

	name, err := valueobject.NewUserName(req.GetName())
	if err != nil {
		return &authpb.RegisterUserResponse{Ok: false, Error: err.Error()}, nil
	}

	email, err := valueobject.NewEmail(req.GetEmail())
	if err != nil {
		return &authpb.RegisterUserResponse{Ok: false, Error: err.Error()}, nil
	}

	cmd := dto.RegisterUserCommand{
		Name:     name,
		Email:    email,
		Password: req.GetPassword(),
	}

	result, err := s.registerUser.Execute(cmd)
	if err != nil {
		return &authpb.RegisterUserResponse{Ok: false, Error: err.Error()}, nil
	}

	return &authpb.RegisterUserResponse{
		Ok:     true,
		UserId: result.UserID.Value(),
		Name:   result.Name.Value(),
	}, nil
}

func (s *AuthServer) LoginUser(
	ctx context.Context,
	req *authpb.LoginUserRequest,
) (*authpb.LoginUserResponse, error) {

	email, err := valueobject.NewEmail(req.GetEmail())
	if err != nil {
		return &authpb.LoginUserResponse{Ok: false, Error: err.Error()}, nil
	}

	cmd := dto.LoginUserCommand{
		Email:    email,
		Password: req.GetPassword(),
	}

	result, err := s.loginUser.Execute(cmd)
	if err != nil {
		return &authpb.LoginUserResponse{Ok: false, Error: err.Error()}, nil
	}

	principal, err := identity.NewPrincipal(
		result.UserID,
		[]identity.Role{identity.RoleUser},
		"auth-service",
		time.Now(),
		time.Now().Add(24*time.Hour),
	)
	if err != nil {
		return &authpb.LoginUserResponse{Ok: false, Error: "token issuance failed"}, nil
	}

	token, err := s.tokenIssuer.IssueToken(principal)
	if err != nil {
		return &authpb.LoginUserResponse{Ok: false, Error: "token issuance failed"}, nil
	}

	return &authpb.LoginUserResponse{
		Ok:     true,
		UserId: result.UserID.Value(),
		Name:   result.Name.Value(),
		Token:  token,
	}, nil
}
