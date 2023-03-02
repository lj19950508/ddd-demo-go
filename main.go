package main

import (
	"go.uber.org/fx"
)


func main() {
	fx := fx.New(
		option()...,
	)
	fx.Run()
}



