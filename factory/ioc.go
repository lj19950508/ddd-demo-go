package factory

import (
	"fmt"
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

func RegisterAllToIOC(ioc IOC,data map[reflect.Type]any) {
	ioc.data = data
}

func RegisterToIOC(ioc *IOC,instance any) {
	fieldType := reflect.TypeOf(instance)
	ioc.data[fieldType] = instance
}

//传入类型与传出一致
func GetFromIOC[T any](ioc IOC,instance T) T {
	fieldType := reflect.TypeOf(instance)
	fmt.Println(fieldType)
	fmt.Println(ioc.data)
	return ioc.data[fieldType].(T)
}
