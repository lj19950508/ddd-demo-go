package query


//查询结构层

type PageResult[T any] struct{
	List []T `json:"list"`
	Total int64 `json:"total"`
}

func NewPageResult[T any](list []T,count int64) *PageResult[T]{
	return &PageResult[T]{list,count}
}



type UserResult  struct {
	ID *int  `json:"id"`
	Name *string `json:"name"`
}

func NewUserResult(id *int,Name *string) *UserResult{
	return &UserResult{id,Name}
}