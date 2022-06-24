package gorm

import repository "ddd-demo-go/domain/biz1/repository"
import entity "ddd-demo-go/domain/biz1/entity"

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() repository.UserRepository {
	return UserRepositoryImpl{}
}

func (t UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	return nil, nil
}

func (t UserRepositoryImpl) Save(user entity.User) {

}
