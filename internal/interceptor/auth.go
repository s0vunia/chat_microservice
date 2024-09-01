package interceptor

import (
	"context"

	"github.com/s0vunia/chat_microservice/internal/client/authservice"
	"github.com/s0vunia/chat_microservice/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthInterceptor interceptor
func AuthInterceptor(client authservice.AuthService) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Info(
			"handling request",
			zap.String("method", info.FullMethod),
		)
		md, _ := metadata.FromIncomingContext(ctx)
		newCtx := metadata.NewOutgoingContext(ctx, md)
		err := client.Check(newCtx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}
