package query

import "database/sql"

//查询结构层

type UserResult  struct {
	ID sql.NullInt64  `json:"id"`
	Name sql.NullString `json:"name"`
}

func NewUserResult(id sql.NullInt64,Name sql.NullString) *UserResult{
	return &UserResult{id,Name}
}