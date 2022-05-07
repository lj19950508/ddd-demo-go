package gorm

import repository "ddd-demo-go/domain/biz1/repository"
import entity "ddd-demo-go/domain/biz1/entity"

type UserRepositoryImpl struct {
}

func New() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (this UserRepositoryImpl) FindById(id uint) (*entity.User, error) {
	return nil, nil
}

func (this UserRepositoryImpl) Save(user entity.User) {}
