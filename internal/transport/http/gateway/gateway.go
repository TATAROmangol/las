package gateway

import (
	"context"
	test "lms/pkg/api/test/api"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	server *http.Server
}

func NewServer(cfg Config, ctx context.Context) (*Gateway, error) {
	con, err := grpc.NewClient(
		cfg.GrpcPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	gwmux := runtime.NewServeMux()
	err = test.RegisterOrderServiceHandler(ctx, gwmux, con)
	if err != nil {
		return nil, err
	}

	gwServer := &http.Server{
		Addr:    cfg.GWPort,
		Handler: gwmux,
	}

	return &Gateway{gwServer}, nil
}

func (g *Gateway) Run() error {
	if err := g.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (g *Gateway) GracefulStop(ctx context.Context) error {
	if err := g.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
