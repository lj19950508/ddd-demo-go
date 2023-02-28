package ioc

import (
	"fmt"
	"reflect"
)

var ioc = make(map[reflect.Type]any)

func RegisterAll(data []any) {
	for i := 0; i < len(data); i++ {
		Register(data[i])
	}
}

func Register(instance any) {
	//instance必须为指针
	if reflect.ValueOf(instance).Kind() != reflect.Pointer {
		panic("Cloud not register an object without pointer")
	}

	instanceType := reflect.TypeOf(instance)

	if ioc[instanceType] != nil {
		panic("Could not register an exists object")
	}
	ioc[instanceType] = instance
}

func Get[T any]() *T {
	//(*T)(nil) 声明一个类型T为空 约等于 c#的 typeof(T)
	instanceType := reflect.TypeOf((*T)(nil))
	//try error
	val, ok := ioc[instanceType].(*T)
	if !ok {
		panic(fmt.Sprintf("Could not find in ioc: %s", instanceType))
	}
	return val
}
