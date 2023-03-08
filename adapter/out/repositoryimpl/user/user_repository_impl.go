package repositoryimpl

import (
	"github.com/jinzhu/gorm"
	"github.com/lj19950508/ddd-demo-go/adapter/out/repositoryimpl/user/po"
	user "github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/pkg/errors"

	// "github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
)

type UserRepositoryImpl struct {
	*db.DB
	logger.Interface
}

//------------------

func NewUserRepositoryImpl(mysql *db.DB, logger logger.Interface) user.UserRepository {
	return &UserRepositoryImpl{
		mysql,
		logger,
	}
}

func (t *UserRepositoryImpl) Load(id int64) (*user.User, error) {
	//获取单挑po、
	var userPo po.User
	if result := t.GormDb.First(&userPo, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, user.ErrUserNoExists
		}
		return nil, errors.WithStack(result.Error)
	}
	//把po->domain
	domainUser := user.NewUser(userPo.ID, userPo.Name)
	return domainUser, nil
}

func (t *UserRepositoryImpl) Add(user *user.User) error {
	//更具IDsave
	result := t.GormDb.Create(&po.User{
		Name: user.Name,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *UserRepositoryImpl) Save(user *user.User) error {
	userPo := po.User{
		ID:   user.Id,
		Name: user.Name,
	}
	result := t.GormDb.Save(&userPo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		t.Error("0 rows affected when save,%+v", &userPo)
	}
	return nil
}

func (t *UserRepositoryImpl) Remove(user *user.User) error {
	//更具IDsave
	result := t.GormDb.Delete(&po.User{
		ID:   user.Id,
		Name: user.Name,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
