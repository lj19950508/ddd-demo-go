package factory

import (
	"fmt"
	"testing"
)

type Test struct {
}
type Test1 struct {
	a int
}

func TestIOC(t *testing.T) {

	// var list=
	RegisterAll([]any{
		&Test1{a: 1},
		&Test{},
	})
	// Register(&Test1{
	// 	a: 1,
	// })
	a := Get[Test]()
	fmt.Println(a)

	// ioc := NewIOC()
	// var a = Test1{}
	// Register(&Test1{
	// 	a: 1,
	// })
	// a := Get[Test1]()
	// // a = Get[Test1]()
	// fmt.Println(a)
	// a=ioc.Get(Test1{})

	// RegisterToIOC(ioc, Test1{})
	// // Register(&Test1{})
	// // Register(&Test2{})

	// // var a *Test1
	// // var aa *Test2
	// // var b = (*Test1)(nil)
	// fmt.Println(Test1{})
	// fmt.Println(unsafe.Sizeof(Test1{}))
	// // var b = (*Test1)(nil)
	// fmt.Println(unsafe.Sizeof((*Test1)(nil)))

	// //传入了一个空类型
	// a := GetFromIOC(*ioc, Test1{})
	// fmt.Println(a)
	// Get[Test]()
	// test :=Test{}

	// test = Get(Test.(type))
}
