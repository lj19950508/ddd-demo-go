package domain

type User struct {
	Id   uint
	Name string
}

func NewUser(Id uint, name string) *User {
	return &User{Id: Id, Name: name}
}

//具体业务数据操作
