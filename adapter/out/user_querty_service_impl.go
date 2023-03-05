package out

import "github.com/lj19950508/ddd-demo-go/application/query"

// TODO 这个service可以从各种地方 es,db,redis,混合查询 

type UserQueryServiceImpl struct{
	//es1 db redis
}

func NewUserQueryServiceimpl() *UserQueryServiceImpl{
	return &UserQueryServiceImpl{}
}

func (t *UserQueryServiceImpl) Info(query query.UserQuery) (*query.UserResult, error) {
	return nil,nil
	//1.组建一个事务甚至是分布式事务
	//2. 操作仓储或者队列
	//3. 操作domain,或者领域服务
	//4. 返回一个domain 
	// user,err := t.userRepository.Load(id)
	// if(user==nil){
	// 	return nil,userpkg.ErrOrderStatusError
	// }
	// return user,err
}
