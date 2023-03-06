package repositoryimpl

import (
	"database/sql"

	user "github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/pkg/errors"

	// "github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"gorm.io/gorm"
)

type UserPO struct {
	//可空得用指针
	//非空用int
	//ID主键
	ID   sql.NullInt64
	Name sql.NullString
}

func NewUserPO(id int64, name string) *UserPO {
	return &UserPO{
		ID:   sql.NullInt64{Int64: id,Valid: true},
		Name: sql.NullString{String: name,Valid: true},
	}
}

//---------------------

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
	t.GormDb.Rollback()
	var userPo UserPO
	if result := t.GormDb.First(&userPo, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(result.Error)
	}
	//把po->domain
	domainUser := user.NewUser(userPo.ID.Int64, userPo.Name.String)
	return domainUser, nil
}

func (t *UserRepositoryImpl) Save(user *user.User) error {
	//Fetch save 模型  Update都必须取回才能操作
	//do -> po
	userPo := NewUserPO(5, "test")
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
	//这里可以把userpo和user何在一起吗？
	//Fetch save 模型  Update都必须取回才能操作
	//do -> po
	userPo := NewUserPO(5, "test")
	//更具IDsave
	result := t.GormDb.Create(userPo)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
