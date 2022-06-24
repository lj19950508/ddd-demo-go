package domain

import entity "ddd-demo-go/domain/biz1/entity"

type UserRepository interface {

	//domain
	FindById(id int) (*entity.User, error)

	//FindList() []*entity.User

	Save(user entity.User)
}
