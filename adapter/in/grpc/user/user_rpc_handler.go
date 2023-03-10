package grpchandler

import (
	"context"

	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg/bizerror"
	"github.com/lj19950508/ddd-demo-go/protos/user"
	"google.golang.org/grpc"
)

type UserApi struct {
	pb.UnimplementedUserCenterServer
}

func (s *UserApi) Register(grpc *grpc.Server) {
	pb.RegisterUserCenterServer(grpc, s)
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

func (s *UserApi) Login(ctx context.Context, in *pb.SaveEvent) (*pb.SaveEventRes, error) {
	//加入这里是另一个服务端
	return &pb.SaveEventRes{Msg: "Hello " + string(rune(in.Id))}, nil
	//eventbus.pu
}
