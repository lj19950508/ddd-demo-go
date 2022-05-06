package repository

type UserRepository interface {
	//domain
	FindById(id uint) (*entity.User, error)

	//FindList() []*entity.User

	Save(user entity.User)
}
