package kernel

import (
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/contract"
	"github.com/theone-daxia/bdj/framework/gin"
)

type BdjKernelProvider struct {
	HttpEngine *gin.Engine
}

func (p *BdjKernelProvider) Name() string {
	return contract.KernelKey
}

func (p *BdjKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewBdjKernelService
}

func (p *BdjKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{p.HttpEngine}
}

func (p *BdjKernelProvider) IsDefer() bool {
	return false
}

func (p *BdjKernelProvider) Boot(c framework.Container) error {
	if p.HttpEngine == nil {
		p.HttpEngine = gin.Default()
	}
	p.HttpEngine.SetContainer(c)
	return nil
}
