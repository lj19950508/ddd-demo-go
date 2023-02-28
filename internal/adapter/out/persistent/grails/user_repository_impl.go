package grails

import (
	"github.com/lj19950508/ddd-demo-go/internal/adapter/out/persistent/grails/pojo"
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
	//获取单挑po 
	userPo := pojo.NewUserPO()

	result := t.GormDb.First(&userPo, id)
	//处理数据库异常
	if result.Error != nil {
		return nil, result.Error
	}
	//把po->domain
	domainUser := entity.NewUser(userPo.Id, userPo.Name)
	return domainUser, nil
}

func (t *UserRepositoryImpl) Save(user entity.User) {

}
