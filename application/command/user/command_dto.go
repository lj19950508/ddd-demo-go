package command



type UpdateCommand struct{
	ID int64 `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}
type CreateCommand struct{
	Name string `form:"name" json:"name"`
}