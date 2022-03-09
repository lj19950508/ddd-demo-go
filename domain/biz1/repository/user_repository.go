package repository

import "ddd-demo1/domain/biz1/entity"

type UserRepository interface {
	//domain
	FindById(id uint64) *entity.User

	//FindList() []*entity.User

	Save(user entity.User)
}
