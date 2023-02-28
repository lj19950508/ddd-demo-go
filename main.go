package main

import (
	"log"

	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/internal/app"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		//这么打才能打印出堆栈信息
		// %v在打印接口类型时，会打印出其实际的值。而在打印结构体对象时，打印的是结构体成员对象的值。
		// %+v打印结构体对象中的字段类型+字段值。
		// %#v先打印结构体名，再输出结构体对象的字段类型+字段的值。
		log.Fatalf("%+v", err)
		//这个就是 log+ exits  也就是 panic的意思 ， 但是panic会被 recover 
	}
	app.Run(cfg)
}
