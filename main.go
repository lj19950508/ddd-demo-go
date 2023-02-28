package main

import (
	"github.com/lj19950508/ddd-demo-go/config"
	"github.com/lj19950508/ddd-demo-go/internal/app"
)
func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		//配置文件是最底层的,所以直接panic
		panic(err)
	}

	app.Run(cfg)
}

