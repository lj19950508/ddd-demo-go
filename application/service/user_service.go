package service

type UserService struct {
}

func NewUserService() IUserService {
	return &UserService{}
}

func (this *UserService) Hello() {
	println("hello world")

}
