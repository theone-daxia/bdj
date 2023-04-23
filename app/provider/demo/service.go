package demo

import (
	"fmt"
	"github.com/theone-daxia/bdj/framework"
)

type Service struct {
	c framework.Container
}

func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "student 1",
		},
		{
			ID:   2,
			Name: "student 2",
		},
	}
}

func NewService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	fmt.Println("new demo service")
	return &Service{c: c}, nil
}
