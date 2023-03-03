package user

import (
	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg/bizerror"
)

type User struct {
	Id   uint
	Name string
}

func NewUser(Id uint, name string) *User {
	return &User{Id: Id, Name: name}
}

//具体业务数据操作
//领域服务

type UserRepository interface {

	//domain
	FindById(id int) (*User, error)

	//FindList() []*entity.User

	Save(user *User) error
}

var (
	//业务异常码
	ErrOrderStatusError = bizerror.NewBizError(100,"错误错误")
)