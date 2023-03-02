package grails

import (
	"github.com/lj19950508/ddd-demo-go/internal/adapter/out/persistent/grails/po"
	entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"
	repository "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/repository"
	"github.com/pkg/errors"
	// "github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/mysql"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*mysql.Mysql
	logger.Interface
}

func NewUserRepositoryImpl(mysql *mysql.Mysql,logger logger.Interface) repository.UserRepository {
	return &UserRepositoryImpl{
		mysql,
		logger,
	}
}

func (t *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	
	//获取单挑po、
	var userPo po.User
	if result := t.GormDb.First(&userPo, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(result.Error)
	}
	//把po->domain
	domainUser := entity.NewUser(userPo.ID, userPo.Name)
	return domainUser, nil
}

//save -> void cqrs有点
func (t *UserRepositoryImpl) Save(user *entity.User) error {
	//do -> po
	userPo := po.NewUserPO(5, "test")

	if userPo.ID == 0 {
		//需要返回user则加&
		if result := t.GormDb.Create(userPo); result.Error != nil {
			return result.Error
		}
	} else {
		result := t.GormDb.Model(userPo).Updates(userPo)
		if result.RowsAffected == 0 {
			t.Warn("0 rows affected,%+v", userPo)
		}
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
