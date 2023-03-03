package service

import (
	userpkg "github.com/lj19950508/ddd-demo-go/domain/user"
)

type UserService interface {
	Info(id int) (*userpkg.User, error)
}

type UserServiceImpl struct {
	userRepository userpkg.UserRepository
}

func NewUserServiceImpl(userRepository userpkg.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (t *UserServiceImpl) Info(id int) (*userpkg.User, error) {
	//1.组建一个事务甚至是分布式事务
	//2. 操作仓储或者队列
	//3. 操作domain,或者领域服务
	//4. 返回一个domain 
	user,err := t.userRepository.FindById(id)
	if(user==nil){
		return nil,userpkg.ErrOrderStatusError
	}
	return user,err
}
