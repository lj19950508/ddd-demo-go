package queryimpl

import "github.com/lj19950508/ddd-demo-go/application/query"

// TODO 这个service可以从各种地方 es,db,redis,混合查询

type UserQueryServiceImpl struct {
	query.UserQueryService
	//es1 db redis
}

func NewUserQueryServiceimpl() *UserQueryServiceImpl {
	return &UserQueryServiceImpl{}
}

func (t *UserQueryServiceImpl) FindOne(query *query.UserQuery) (*query.UserResult, error) {
	return nil, nil

}

func (t *UserQueryServiceImpl) FindList(query *query.UserQuery) ([]query.UserResult, error) {
	return nil, nil
}
