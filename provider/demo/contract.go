package demo

const Key = "bdj:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
