package command

import (
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
	err := t.eventBus.Publish(eventbus.NewEvent(0,userpkg.EvtUserCreate,nil))
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
