package grpcv1

import (
	"context"
	"lms/pkg/logger"
	"log/slog"

	"google.golang.org/grpc"
)

func LoggerInterceptor(log *slog.Logger) func(context.Context, any, *grpc.UnaryServerInfo, grpc.UnaryHandler,) (any, error){
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = logger.InitToCtx(ctx, log)

		logger.GetFromCtx(ctx).InfoContext(ctx, "grpc server call", "method", info.FullMethod)

		resp, err := handler(ctx, req)
		if err != nil{
			logger.GetFromCtx(ctx).ErrorContext(ctx, "failed in handler", "error", err)
		}

		return resp, nil
	}
}