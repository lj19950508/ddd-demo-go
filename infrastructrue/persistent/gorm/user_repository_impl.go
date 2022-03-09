package gorm

import (
	"ddd-demo1/domain/biz1/entity"
	"ddd-demo1/domain/biz1/repository"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (this UserRepositoryImpl) FindById(id uint64) *entity.User {
	return &entity.User{}
}

//func (this UserRepositoryImpl) FindList() []*entity.User{
//	return []*entity.User{}
//}

func (this UserRepositoryImpl) Save(user entity.User) {}
