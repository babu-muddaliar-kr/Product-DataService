package main

import (
	"Product-DataService/handlers"
	"github.com/krogertechnology/krogo/pkg/krogo"
)

func main() {
	k := krogo.New()
	k.Server.ValidateHeaders = false
	k.Server.Router.Prefix("/product-DataService")
	k.GET("/greet", handlers.Greet)
	k.Start()
}
