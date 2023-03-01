package po

//表名为
type User struct {
	//可空得用指针
	//非空用int
	//ID主键
	ID   uint
	Name string
}

func NewUserPO() User {
	return User{}
}
