package po

type User struct {
	//可空得用指针
	//非空用int
	//ID主键
	ID   int64
	Name string
}

func NewUserPO(id int64, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

//---------------------