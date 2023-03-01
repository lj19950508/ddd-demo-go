package domain

import entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"

type UserRepository interface {

	//domain
	FindById(id int) (*entity.User, error)

	//FindList() []*entity.User
	
	Save(user *entity.User) error
}
