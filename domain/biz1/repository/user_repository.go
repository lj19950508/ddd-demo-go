package repository

import "ddd-demo1/domain/biz1/entity"

type UserRepository interface {
	//domain
	FindById(id uint) (*entity.User, error)

	//FindList() []*entity.User

	Save(user entity.User)
}
