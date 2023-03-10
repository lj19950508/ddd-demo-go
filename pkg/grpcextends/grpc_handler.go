package grpcextends

import "google.golang.org/grpc"

type GrpcHandler interface {
	Register(s *grpc.Server)
}
