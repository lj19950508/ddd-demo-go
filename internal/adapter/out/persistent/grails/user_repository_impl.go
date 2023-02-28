package grails

import (
	"github.com/lj19950508/ddd-demo-go/internal/adapter/out/persistent/grails/pojo"
	entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"
	repository "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/repository"
	"github.com/lj19950508/ddd-demo-go/pkg/mysql"
	"github.com/pkg/errors"
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
	userPo := pojo.NewUserPO()

	result := t.GormDb.First(&userPo, id)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	domainUser := entity.NewUser(userPo.Id, userPo.Name)
	// copier.Copy()
	//拷贝userpo成entity.user
	return domainUser, nil
}

func (t *UserRepositoryImpl) Save(user entity.User) {

}
