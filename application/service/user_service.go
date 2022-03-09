package service

import "ddd-demo1/domain/biz1/entity"

type UserService interface {
	Info(id uint) (*entity.User, error)
}
