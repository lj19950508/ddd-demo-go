package po

type User struct {
	//可空得用指针
	//非空用int
	//ID主键
	ID   int64
	Name string
}
