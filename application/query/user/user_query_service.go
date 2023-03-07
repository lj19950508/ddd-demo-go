package query

type UserQueryService interface {

	FindOne(query *UserQuery) (*UserResult,error)

	FindList(query *UserPageQuery) (*PageResult[UserResult],error)

}

