package gorm

import (
	"ddd-demo1/domain/biz1/entity"
	"ddd-demo1/domain/biz1/repository"
	"ddd-demo1/infrastructrue/middleware"
	"ddd-demo1/infrastructrue/persistent/gorm/pojo"
)

type UserRepositoryImpl struct {
	resource *middleware.GormResource
}

func NewUserRepositoryImpl(resource *middleware.GormResource) repository.UserRepository {
	return &UserRepositoryImpl{
		resource: resource,
	}
}

func (this UserRepositoryImpl) FindById(id uint) (*entity.User, error) {
	var user pojo.User
	result := this.resource.DB().First(&user, id)
	return entity.NewUser(
		user.ID,
		user.Name,
	), result.Error
}

func (this UserRepositoryImpl) Save(user entity.User) {}
