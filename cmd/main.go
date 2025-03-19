package main

import (
	"context"
	"lms/internal/config"
	"lms/pkg/logger"
	"lms/internal/service"
	"lms/internal/storage"
	grpcv1 "lms/internal/transport/grpcv1"
	"lms/internal/transport/http/gateway"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	ctx := context.Background()

	log := logger.New()

	repo := storage.NewOrderRepo()
	service := service.NewOrderService(repo)
	grpcServer := grpcv1.New(cfg.GrpcConfig, service, log)

	go func() {
		if err := grpcServer.Run(); err != nil {
			log.ErrorContext(ctx, "failed run grpc server", "error", err)
			os.Exit(1)
		}
	}()

	log.Info("grpc server started", "address", cfg.GrpcConfig.GrpcPort)

	gwServer, err := gateway.NewServer(cfg.GWConfig, ctx)
	if err != nil {
		log.Error("failed create gateway server", "error", err)
		os.Exit(1)
	}

	go func() {
		if err := gwServer.Run(); err != nil && err != http.ErrServerClosed {
			log.Error("failed run gateway server", "error", err)
			os.Exit(1)
		}
	}()
	log.Info("gateway server started", "address", cfg.GWConfig.GWPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("start shutdown")

	closeCtx, closeCancel := context.WithTimeout(ctx, 5*time.Second)
	defer closeCancel()

	if err := gwServer.GracefulStop(closeCtx); err != nil {
		log.Error("failed shutdown gateway", "error", err)
		os.Exit(1)
	}

	grpcServer.GracefulStop()

	log.Info("Server Stopped")
}
