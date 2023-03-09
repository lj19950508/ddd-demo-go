package command

import (
	"fmt"

	userpkg "github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/lj19950508/ddd-demo-go/pkg/eventbus"
)

type UserCommandService interface {
	Create(cmd *CreateCommand) error
	Update(cmd *UpdateCommand) error
	Delete(id int64) error
}

//---------------------------

type UserCommandImpl struct {
	userRepository userpkg.UserRepository
	eventBus       eventbus.EventBus
}

func NewUserCommandImpl(userRepository userpkg.UserRepository, eventBus eventbus.EventBus) UserCommandService {
	return &UserCommandImpl{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

func (t UserCommandImpl) Create(cmd *CreateCommand) error {
	user := userpkg.NewUser(0, cmd.Name)
	if err := t.userRepository.Add(user); err != nil {
		return err
	}
	evt:=eventbus.NewEvent(10,userpkg.EvtUserCreate,"string")
	evt.Return=true
	evt.Compensation="UserBuChang"
	err := t.eventBus.Publish(evt)
	fmt.Printf("收到了回执%+v\n",evt.ExcuteResult)
	if err != nil {
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
