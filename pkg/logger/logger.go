package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	Key = "logger"
)

func New() *slog.Logger{
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
}

func InitToCtx(ctx context.Context, log *slog.Logger) context.Context{
	return context.WithValue(ctx, Key, log)
}

func GetFromCtx(ctx context.Context) *slog.Logger{
	return ctx.Value(Key).(*slog.Logger)
}

