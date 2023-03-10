package grpcextends

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcHandler interface {
	Register(s *grpc.Server)
}

func Send[T any](url string, cb func(*grpc.ClientConn, context.Context) (T, error)) (result T,err  error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res,err:= cb(conn, ctx)
	defer cancel()
	return res,err
}
