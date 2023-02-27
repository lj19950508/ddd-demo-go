package service

import (
	entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"
	repository "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/repository"
)

type UserService interface {
	Info(id int) (*entity.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (t *UserServiceImpl) Info(id int) (*entity.User, error) {
	return t.userRepository.FindById(id)
}
