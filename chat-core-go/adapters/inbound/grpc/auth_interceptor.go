package grpcadapter

import (
	"chat-core-go/adapters/outbound/jwtadapter"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const userIDKey contextKey = "userId"

func AuthInterceptor(jwtService *jwtadapter.JWTService) grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {

        md, ok := metadata.FromIncomingContext(ctx)
        if !ok {
            return nil, status.Error(codes.Unauthenticated, "missing metadata")
        }

        values := md.Get("authorization")
        if len(values) == 0 {
            return nil, status.Error(codes.Unauthenticated, "missing authorization header")
        }

        token := values[0]

        userId, err := jwtService.VerifyToken(token)
        if err != nil {
            return nil, status.Error(codes.Unauthenticated, "invalid token")
        }

        // put userId into context
        ctx = context.WithValue(ctx, userIDKey, userId)

        // continue to real handler
        return handler(ctx, req)
    }
}
