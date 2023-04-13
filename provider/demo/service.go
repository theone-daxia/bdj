package demo

import (
	"fmt"
	"github.com/theone-daxia/bdj/framework"
)

type DemoService struct {
	Service // 实现接口
	c       framework.Container
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	fmt.Println("new demo service")
	return &DemoService{c: c}, nil
}
