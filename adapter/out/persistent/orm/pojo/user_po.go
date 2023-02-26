package pojo

//表名为
type User struct {
	//ID主键
	ID   int
	Name string
}

func NewUserPo() *User {
	return &User{}
}
