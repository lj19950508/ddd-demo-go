package command

import (
	userpkg "github.com/lj19950508/ddd-demo-go/domain/user"
)

type UserCommandService interface {
	Save(cmd SaveCommand) error
}

//---------------------------

type UserCommandImpl struct {
	userRepository userpkg.UserRepository
}

func NewUserCommandImpl(userRepository userpkg.UserRepository) UserCommandService {
	return &UserCommandImpl{
		userRepository: userRepository,
	}
}

func (t UserCommandImpl) Save(cmd SaveCommand) error {
	return nil
}
