package grails

import (
	entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"
	repository "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db:db,
	}
}

func (t *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	return nil, nil
}

func (t *UserRepositoryImpl) Save(user entity.User) {

}
