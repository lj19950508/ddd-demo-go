package rpcsender

import "github.com/lj19950508/ddd-demo-go/protos/user"

type RpcSender interface {
	ProductSave(req *pb.SaveEvent) (*pb.SaveEventRes, error)
}
