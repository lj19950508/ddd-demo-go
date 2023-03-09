package user

import (
	"context"

	"github.com/lj19950508/ddd-demo-go/protos/user"
	"google.golang.org/grpc"
)

type UserApi struct{
	pb.UnimplementedUserCenterServer
}

func NewUserApi (s *grpc.Server) *UserApi{
	UserApi := &UserApi{}
	pb.RegisterUserCenterServer(s,UserApi)
	return  UserApi
}

func (s *UserApi) Login(ctx context.Context, in *pb.SaveEvent) (*pb.SaveEventRes, error) {
	return &pb.SaveEventRes{Msg: "Hello " + string(rune(in.Id))}, nil
}
