package service

import (
	"ddd-demo1/domain/biz1/entity"
	"ddd-demo1/domain/biz1/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (this *UserServiceImpl) Hello() *entity.User {
	return this.userRepository.FindById(1)
}
