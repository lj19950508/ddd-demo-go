package entity

type User struct {
	id   uint
	name string
}

func NewUser(id uint, name string) *User {
	return &User{id: id, name: name}
}
