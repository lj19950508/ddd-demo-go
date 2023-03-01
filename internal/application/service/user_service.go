package service

import (
	entity "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/entity"
	repository "github.com/lj19950508/ddd-demo-go/internal/domain/biz1/repository"
	"github.com/lj19950508/ddd-demo-go/pkg"
)

type UserService interface {
	Info(id int) (*entity.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (t *UserServiceImpl) Info(id int) (*entity.User, error) {
	//1.组建一个事务甚至是分布式事务
	//2. 操作仓储或者队列
	//3. 操作domain,或者领域服务
	//4. 返回一个domain 
	user,err := t.userRepository.FindById(id)
	if(user==nil){
		return nil,pkg.ErrOrderStatusError
	}
	return user,err
}
