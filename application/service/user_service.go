package service

import "ddd-demo1/domain/biz1/entity"

type UserService interface {
	Hello() *entity.User
}
