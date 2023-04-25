package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会替换，不返回 error。
	Bind(provider ServiceProvider) error

	// IsBind 关键字凭证是否已经绑定服务提供者。
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务。
	Make(key string) (interface{}, error)

	// MakeNew 根据关键字凭证和参数获取一个新的服务。
	MakeNew(key string, params []interface{}) (interface{}, error)

	// MustMake 根据关键字凭证获取一个服务。
	// 如果关键字凭证未绑定服务提供者，会 panic。
	// 所以在使用本接口时，要保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
}

type BdjContainer struct {
	// 强制要求 BdjContainer 实现 Container 接口
	Container
	// 存储注册的服务提供者
	providers map[string]ServiceProvider
	// 存储具体的实例
	instances map[string]interface{}
	// 锁住对容器的变更操作
	lock sync.RWMutex
}

func NewBdjContainer() *BdjContainer {
	return &BdjContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (b *BdjContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range b.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

func (b *BdjContainer) Bind(provider ServiceProvider) error {
	b.lock.Lock()
	key := provider.Name()

	b.providers[key] = provider
	b.lock.Unlock()

	// 推迟实例化，则直接返回
	if provider.IsDefer() {
		return nil
	}

	ins, err := b.newInstance(provider, nil)
	if err != nil {
		return err
	}
	b.instances[key] = ins
	return nil
}

func (b *BdjContainer) IsBind(key string) bool {
	return b.findServiceProvider(key) != nil
}

func (b *BdjContainer) findServiceProvider(key string) ServiceProvider {
	b.lock.RLock()
	defer b.lock.RUnlock()
	if sp, ok := b.providers[key]; ok {
		return sp
	}
	return nil
}

// Make 调用 make 实例化服务
func (b *BdjContainer) Make(key string) (interface{}, error) {
	return b.make(key, nil, false)
}

// MakeNew 调用 make 强制重新实例化服务
func (b *BdjContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return b.make(key, params, true)
}

// MustMake 调用 make 实例化服务，有错误会 panic
func (b *BdjContainer) MustMake(key string) interface{} {
	ins, err := b.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return ins
}

// make 真正实例化服务
func (b *BdjContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	sp := b.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("provider:" + key + " is not registered")
	}

	// 需要强制重新实例化
	if forceNew {
		return b.newInstance(sp, params)
	}

	// 服务器容器中已经存在对应的服务实例，直接返回
	if ins, ok := b.instances[key]; ok {
		return ins, nil
	}

	ins, err := b.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	b.instances[key] = ins
	return ins, nil
}

func (b *BdjContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(b); err != nil {
		return nil, err
	}

	if params == nil {
		params = sp.Params(b)
	}
	method := sp.Register(b)
	ins, err := method(params...)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
