package domain

type User struct {
	Id   int
	Name string
}

func NewUser(Id int, name string) *User {
	return &User{Id: Id, Name: name}
}

//具体业务数据操作
