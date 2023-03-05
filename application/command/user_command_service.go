package command

import (
	userpkg "github.com/lj19950508/ddd-demo-go/domain/user"
)

type UserCommandService interface {
	Save(cmd SaveCommand) error
}

//---------------------------

type UserCommandImpl struct {
	userRepository userpkg.UserRepository
}

func NewUserCommandImpl(userRepository userpkg.UserRepository) UserCommandService {
	return &UserCommandImpl{
		userRepository: userRepository,
	}
}

func (t UserCommandImpl) Save(cmd SaveCommand) error {
		//1.组建一个事务甚至是分布式事务
	//2. 操作仓储或者队列
	//3. 操作domain,或者领域服务
	//4. 返回一个domain 
	// user,err := t.userRepository.Load(id)
	// if(user==nil){
	// 	return nil,userpkg.ErrOrderStatusError
	// }
	// return user,err
	return nil
}
