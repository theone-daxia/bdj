package framework

type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider 服务提供者需要实现的接口
type ServiceProvider interface {
	// Name 获取服务凭证
	Name() string

	// Register 服务实例化
	Register(Container) NewInstance

	// Params 获取服务实例化的参数
	Params(Container) []interface{}

	// IsDefer 控制服务实例化的时机
	// false 表示不延迟实例化（在注册的时候就实例化）
	// true 表示延迟实例化（在获取服务的时候实例化）
	IsDefer() bool

	// Boot 服务实例化的时候会调用，做一些准备工作
	// 如果 Boot 返回 error，服务实例化就会失败，返回错误
	Boot(Container) error
}
