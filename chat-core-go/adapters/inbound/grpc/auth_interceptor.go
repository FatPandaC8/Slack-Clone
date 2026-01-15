// adapters/inbound/grpc/auth_interceptor.go
package grpcadapter

import (
	"chat-core-go/domain/identity"
	"chat-core-go/ports/service"
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const principalKey contextKey = "principal"

// AuthInterceptor is a gRPC unary interceptor that authenticates requests
// It uses the Authenticator PORT (polymorphic - works with JWT, OAuth2, etc.)
func AuthInterceptor(authenticator service.Authenticator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		
		// Skip authentication for public endpoints
		if isPublicEndpoint(info.FullMethod) {
			return handler(ctx, req)
		}
		
		// Extract metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}
		
		// Extract authorization header
		values := md.Get("authorization")
		if len(values) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}
		
		// Parse Bearer token
		authHeader := values[0]
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization format")
		}
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		
		// Authenticate using the Authenticator PORT
		credentials := service.BearerCredentials{Token: token}
		principal, err := authenticator.Authenticate(ctx, credentials)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "authentication failed: "+err.Error())
		}
		
		// Inject Principal into context
		ctx = context.WithValue(ctx, principalKey, principal)
		
		// Call the actual handler
		return handler(ctx, req)
	}
}

// GetPrincipal extracts the Principal from context
func GetPrincipal(ctx context.Context) (*identity.Principal, error) {
	principal, ok := ctx.Value(principalKey).(*identity.Principal)
	if !ok || principal == nil {
		return nil, status.Error(codes.Unauthenticated, "no principal in context")
	}
	return principal, nil
}

// isPublicEndpoint checks if an endpoint doesn't require authentication
func isPublicEndpoint(method string) bool {
	publicEndpoints := []string{
		"/chat.ChatService/RegisterUser",
		"/chat.ChatService/LoginUser",
	}
	
	for _, endpoint := range publicEndpoints {
		if method == endpoint {
			return true
		}
	}
	
	return false
}