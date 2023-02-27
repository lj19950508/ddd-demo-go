package grails

import (
	entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"
	repository "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/repository"
	"github.com/lj19950508/ddd-demo-go/pkg/mysql"
)

type UserRepositoryImpl struct {
	*mysql.Mysql
}

func NewUserRepositoryImpl(mysql *mysql.Mysql) repository.UserRepository {
	return &UserRepositoryImpl{
		mysql,
	}
}

func (t *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	return nil, nil
}

func (t *UserRepositoryImpl) Save(user entity.User) {

}
