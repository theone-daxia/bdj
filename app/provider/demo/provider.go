package demo

import (
	"fmt"
	"github.com/theone-daxia/bdj/framework"
)

type DemoServiceProvider struct {
}

func (sp *DemoServiceProvider) Name() string {
	return Key
}

func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

func (sp *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}
