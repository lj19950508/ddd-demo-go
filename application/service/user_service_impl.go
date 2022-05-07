package service

import (
	entity "ddd-demo-go/domain/biz1/entity"
	repository "ddd-demo-go/domain/biz1/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (this *UserServiceImpl) Info(id uint) (*entity.User, error) {
	return this.userRepository.FindById(id)
}
