package service

import (
	entity "ddd-demo-go/domain/biz1/entity"
	repository "ddd-demo-go/domain/biz1/repository"
)

type UserService interface {
	Info(id int) (*entity.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (t *UserServiceImpl) Info(id int) (*entity.User, error) {
	return t.userRepository.FindById(id)
}
