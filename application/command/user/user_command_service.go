package command

import (
	"github.com/lj19950508/ddd-demo-go/application/rpcsender/product"
	userpkg "github.com/lj19950508/ddd-demo-go/domain/user"
	pb "github.com/lj19950508/ddd-demo-go/protos/user"
)

type UserCommandService interface {
	Create(cmd *CreateCommand) error
	Update(cmd *UpdateCommand) error
	Delete(id int64) error
}

//---------------------------

type UserCommandImpl struct {
	userRepository userpkg.UserRepository
	productSender  rpcsender.RpcSender
}

func NewUserCommandImpl(userRepository userpkg.UserRepository,productSender rpcsender.RpcSender) UserCommandService {
	return &UserCommandImpl{
		userRepository: userRepository,
		productSender: productSender,
	}
}

func (t UserCommandImpl) Create(cmd *CreateCommand) error {
	user := userpkg.NewUser(0, cmd.Name)
	if err := t.userRepository.Add(user); err != nil {
		return err
	}
	_,err:=t.productSender.ProductSave(&pb.SaveEvent{Id: 100})
	if(err!=nil){
		return err
	}

	return nil
}

func (t UserCommandImpl) Update(cmd *UpdateCommand) error {
	user, err := t.userRepository.Load(cmd.ID)
	if err != nil {
		return err
	}
	user.Name = cmd.Name
	if err = t.userRepository.Save(user); err != nil {
		return err
	}
	return nil
}

func (t UserCommandImpl) Delete(id int64) error {
	user, err := t.userRepository.Load(id)
	if err != nil {
		return err
	}

	if err := t.userRepository.Remove(user); err != nil {
		return err
	}
	return nil
}
