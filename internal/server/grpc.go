package server

import (
	"time"

	"github.com/oliver258/eagle/pkg/app"
	"github.com/oliver258/eagle/pkg/transport/grpc"
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(cfg *app.ServerConfig) *grpc.Server {

	grpcServer := grpc.NewServer(
		grpc.Network("tcp"),
		grpc.Address(":9090"),
		grpc.Timeout(3*time.Second),
	)

	// register biz service
	// v1.RegisterUserServiceServer(grpcServer, service.Svc.Users())

	return grpcServer
}
