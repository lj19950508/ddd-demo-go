package pojo

//表名为
type User struct {
	//ID主键
	Id   int
	Name string
}

func NewUserPO() *User {
	return &User{}
}
