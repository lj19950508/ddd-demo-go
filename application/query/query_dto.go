package query

import "database/sql"

type UserQuery struct{
	ID sql.NullInt64 `form:"id"`
	Name sql.NullString `form:"name"`

}