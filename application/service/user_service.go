package service

import entity "ddd-demo-go/domain/biz1/entity"

type UserService interface {
	Info(id uint) (*entity.User, error)
}
