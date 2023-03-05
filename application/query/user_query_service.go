package query

type UserQueryService interface {

	FindOne(query *UserQuery) (*UserResult,error)

	FindList(query *UserQuery) ([]UserResult,error)

}

