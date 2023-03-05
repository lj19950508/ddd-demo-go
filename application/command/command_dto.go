package command

import "database/sql"


type SaveCommand struct{
	ID sql.NullInt64 `form:"id" json:"id"`
	Name sql.NullString `form:"name" json:"name"`
}