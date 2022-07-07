package handlers

import "github.com/krogertechnology/krogo/pkg/krogo"

func Greet(c *krogo.Context) (interface{}, error) {
	return "Hello World!", nil
}
