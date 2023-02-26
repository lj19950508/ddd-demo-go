package factory

import (
	"fmt"
	"testing"
)

type Test struct {
}
type Test1 struct {
}
type Test2 struct {
}

func test() {

}

func TestIOC(t *testing.T) {
	ioc := NewIOC()
	RegisterToIOC(ioc, &Test1{})
	// Register(&Test1{})
	// Register(&Test2{})

	var a *Test1
	// var aa *Test2
	// var b = (*Test1)(nil)
	// fmt.Println(unsafe.Sizeof(a))
	// var b = (*Test1)(nil)
	// fmt.Println(unsafe.Sizeof(b))

	//传入了一个空类型
	// ioc :=&IOC[any]{}
	fmt.Println(a)
	a = GetFromIOC(*ioc, (*Test1)(nil))

	// Get[Test]()
	// test :=Test{}

	// test = Get(Test.(type))
}
