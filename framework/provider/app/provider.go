package app

import (
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/contract"
)

type BdjAppProvider struct {
	BaseFolder string
}

func (p *BdjAppProvider) Name() string {
	return contract.AppKey
}

func (p *BdjAppProvider) Register(c framework.Container) framework.NewInstance {
	return NewBdjApp
}

func (p *BdjAppProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c, p.BaseFolder}
}

func (p *BdjAppProvider) IsDefer() bool {
	return false
}

func (p *BdjAppProvider) Boot(c framework.Container) error {
	return nil
}
