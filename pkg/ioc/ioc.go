package ioc

import (
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
		panic("必须是指针才能注册")
	}

	instanceType := reflect.TypeOf(instance)

	if ioc[instanceType] != nil {
		panic("已存在，请勿重复注册")
	}
	ioc[instanceType] = instance
}

func Get[T any]() *T {
	//(*T)(nil) 声明一个类型T为空 约等于 c#的 typeof(T)
	instanceType := reflect.TypeOf((*T)(nil))
	return ioc[instanceType].(*T)
}
