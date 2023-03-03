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

func NewUserPO(ID uint, Name string) *User {
	return &User{
		ID:   ID,
		Name: Name,
	}
}

type UserRepositoryImpl struct {
	*db.DB
	logger.Interface
}

func NewUserRepositoryImpl(mysql *db.DB, logger logger.Interface) user.UserRepository {
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

//save -> void cqrs有点 //save不会忽略空值 所以用create和updates
//db.save 即使是0值也会保存
//db.update  db.update0值不更新  -
//如传入  status 0 时， 会直接重置状态
//由于我们结构体都使用 值而不是指针， 所以默认都是0,所以我们不进行重置 只更新非零值
//所以 数据库中的0值字段都是无意义的，如果create完是0值，则update完 就没办法修改回0值了
//在业务过程中不要使用零值  如状态0 字串空  bool值false(不使用bool值存数据库) float0   金额0
//不太对不太对 要不要经过查询一次再save全部呢 ?
//TODO 想法错了
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
