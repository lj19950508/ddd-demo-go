package grails

import (
	"errors"

	"github.com/lj19950508/ddd-demo-go/internal/adapter/out/persistent/grails/po"
	entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"
	repository "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/repository"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/mysql"
	"gorm.io/gorm"
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
	//获取单挑po、	
	userPo := po.NewUserPO()
	if result := t.GormDb.First(&userPo, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		//意料之中这么写
		panic(result.Error)


		//意料之外这么写
	}
	//把po->domain
	domainUser := entity.NewUser(userPo.ID, userPo.Name)
	return &domainUser, nil
}

func (t *UserRepositoryImpl) Save(user entity.User) {
	//do -> po
	userPo := po.NewUserPO()
	userPo.ID = 5
	userPo.Name = "test"
	if userPo.ID == 0 {
		if result := t.GormDb.Create(&userPo); result.Error != nil {
			panic(result.Error)
		}
	} else {
		result := t.GormDb.Model(&userPo).Updates(userPo)
		if result.RowsAffected == 0 {
			logger.Instance.Warn("0 rows affected,%+v", userPo)
		}
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
