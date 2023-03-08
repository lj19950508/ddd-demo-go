package query


type UserQuery struct{
	IdEq *int `form:"id"`
	NameLike *string `form:"name"`
}

type PageQuery struct{
	Page int `form:"page"`
	Size int `form:"size"  binding:"gt=0,lte=10"`
	//size必传且小雨10
}

type UserPageQuery struct{
	PageQuery
	UserQuery

}