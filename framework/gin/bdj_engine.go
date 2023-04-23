package gin

import "github.com/theone-daxia/bdj/framework"

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}

// engine 实现 container 服务提供者的绑定

func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}
