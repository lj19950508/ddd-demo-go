package query


type UserQuery struct{
	IdEq *int `form:"id"`
	NameLike *string `form:"name"`
}

type PageQuery struct{
	Page *int `form:"page"`
	Size *int `form:"size"`
}

type UserPageQuery struct{
	PageQuery
	UserQuery

}