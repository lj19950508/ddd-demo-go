package infrastructure

import "github.com/golobby/container/v3"

func Produce() {

	//because not ioc, so code in one func , ioc is finding now. golobby
	//out adapter init
	//repo init
	//service init
	//in adapter init\
	var cfg1, err = NewConfig()
	if err != nil {
		panic(err)
	}
	container.Singleton(cfg1)

}
