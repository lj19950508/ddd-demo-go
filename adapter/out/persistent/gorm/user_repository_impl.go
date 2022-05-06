package gorm

type UserRepositoryImpl struct {
}

func New() adapter.UserRepository {
	return &UserRepositoryImpl{}
}

func (this UserRepositoryImpl) FindById(id uint) (*entity.User, error) {

}

func (this UserRepositoryImpl) Save(user entity.User) {}
