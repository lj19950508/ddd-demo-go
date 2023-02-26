package factory

import (
	"reflect"
)



type IOC struct {
	data map[reflect.Type]any
}

func NewIOC() *IOC{
	return &IOC{
		data:make(map[reflect.Type]any),
	}
}

func (s *IOC) RegisterAll(data map[reflect.Type]any) {
	s.data = data
}

func (s IOC) Register(instance any) {
	fieldType := reflect.TypeOf(instance)
	s.data[fieldType] = instance
}

//传入类型与传出一致
func Get[T any](s IOC,instance T) T {
	fieldType := reflect.TypeOf(instance)
	return s.data[fieldType].(T)
}
