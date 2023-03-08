package repositoryimpl

import (
	"github.com/lj19950508/ddd-demo-go/adapter/out/repositoryimpl/user/po"
	user "github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/pkg/errors"

	// "github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"gorm.io/gorm"
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

func (t *UserRepositoryImpl) Load(id int) (*user.User, error) {
	//获取单挑po、
	var userPo po.User
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

func (t *UserRepositoryImpl) Save(user *user.User) error {
	//Fetch save 模型  Update都必须取回才能操作
	//do -> po
	userPo := po.NewUserPO(5, "test")
	//更具IDsave
	result := t.GormDb.Save(userPo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		t.Warn("0 rows affected,%+v", userPo)
	}
	return nil
}

func (t *UserRepositoryImpl) Add(user *user.User) error {
	userPo := po.NewUserPO(5, "test")
	//更具IDsave
	result := t.GormDb.Create(userPo)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


