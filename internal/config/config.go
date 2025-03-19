package config

import (
	"fmt"
	"lms/internal/transport/grpcv1"
	"lms/internal/transport/http/gateway"
	"log"
	"os"
)

type Config struct {
	GrpcConfig grpcv1.Config
	GWConfig   gateway.Config
}

func MustLoad() Config {
	grpcPort, ok := os.LookupEnv("GRPC_PORT")
	if !ok {
		log.Fatal("not find GRPC_PORT")
	}
	grpcAddress := fmt.Sprintf(":%v", grpcPort)

	gatewayPort, ok := os.LookupEnv("GW_PORT")
	if !ok {
		log.Fatal("not find GW_PORT")
	}

	gatewayAddress := fmt.Sprintf(":%v", gatewayPort)

	return Config{
		GrpcConfig: grpcv1.Config{
			GrpcPort: grpcAddress,
		},
		GWConfig: gateway.Config{
			GrpcPort: grpcAddress,
			GWPort:   gatewayAddress,
		},
	}
}
