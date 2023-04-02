package middleware

import (
	"fmt"
	"github.com/theone-daxia/bdj/framework"
)

func Test1() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test1")
		ctx.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test2")
		ctx.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test3")
		ctx.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}
