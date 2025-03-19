package grpcv1

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	cfg    Config
	server *grpc.Server
}

func New(cfg Config, service Servicer, log *slog.Logger) *Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(LoggerInterceptor(log)),
	)

	RegisterService(server, service)

	return &Server{cfg, server}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.cfg.GrpcPort)
	if err != nil {
		return err
	}

	if err := s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}
