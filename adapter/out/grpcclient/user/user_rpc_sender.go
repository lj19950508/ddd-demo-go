package grpcclient

import (
	"context"
	rpcsender "github.com/lj19950508/ddd-demo-go/application/rpcsender/product"
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/pkg/grpcextends"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	pb "github.com/lj19950508/ddd-demo-go/protos/user"
	"google.golang.org/grpc"
)

type UserRpcSender struct {
	target string
	logger logger.Interface
}

func NewUserRpcSender(logger logger.Interface, cfg *config.Config) rpcsender.RpcSender {
	return &UserRpcSender{
		logger: logger,
		target: cfg.GrpcServer.RpcTarget.ProductService,
	}
}

//抽象类

func (s *UserRpcSender) ProductSave(req *pb.SaveEvent) (*pb.SaveEventRes, error) {
	return grpcextends.Send(s.target, func(conn *grpc.ClientConn, ctx context.Context) (*pb.SaveEventRes, error) {
		s.logger.Info("访问到了！")
		c := pb.NewUserCenterClient(conn)
		r, err := c.Login(ctx, req)		
		if err != nil {
			//new biz erro (rpc product save err)
			return nil, err
		}
		return r, nil
	})
}
