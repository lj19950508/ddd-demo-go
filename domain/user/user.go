package user

import (
	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg/bizerror"
)

type User struct {
	Id   int64
	Name string
}

func NewUser(Id int64, name string) *User {
	return &User{Id: Id, Name: name}
}

//具体业务数据操作
//领域服务

type UserRepository interface {

	//domain
	Load(id int) (*User, error)

	//FindList() []*entity.User

	Save(user *User) error
}

var (
	//业务异常码
	errUserDisabled = bizerror.NewBizError(100,"禁用用户")
)


//------------------
type UserService struct{
	//Swap(User1,User2)比如这个
	//多个user domain交互需要用到这个，能用 domain实现则用domain，
	//这里的行为是一个独立的可描述的对象
}


