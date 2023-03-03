package grails

import (
	user "github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/pkg/errors"

	// "github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"gorm.io/gorm"
)

type User struct {
	//可空得用指针
	//非空用int
	//ID主键
	ID   uint
	Name string
}

func NewUserPO(ID uint,Name string) *User {
	return &User{
		ID: ID,
		Name: Name,
	}
}


type UserRepositoryImpl struct {
	*db.DB
	logger.Interface
}

func NewUserRepositoryImpl(mysql *db.DB,logger logger.Interface) user.UserRepository {
	return &UserRepositoryImpl{
		mysql,
		logger,
	}
}

func (t *UserRepositoryImpl) FindById(id int) (*user.User, error) {
	
	//获取单挑po、
	var userPo User
	if result := t.GormDb.First(&userPo, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(result.Error)
	}
	//把po->domain
	domainUser := user.NewUser(userPo.ID, userPo.Name)
	return domainUser, nil
}

//save -> void cqrs有点
func (t *UserRepositoryImpl) Save(user *user.User) error {
	//do -> po
	userPo := NewUserPO(5, "test")

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