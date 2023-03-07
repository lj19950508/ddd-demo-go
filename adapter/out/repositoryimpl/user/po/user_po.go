package po

import "database/sql"
type User struct {
	//可空得用指针
	//非空用int
	//ID主键
	ID   sql.NullInt64
	Name sql.NullString
}

func NewUserPO(id int64, name string) *User {
	return &User{
		ID:   sql.NullInt64{Int64: id,Valid: true},
		Name: sql.NullString{String: name,Valid: true},
	}
}

//---------------------