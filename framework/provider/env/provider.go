package env

import (
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/contract"
)

type BdjEnvProvider struct {
	Folder string
}

func (p *BdjEnvProvider) Name() string {
	return contract.EnvKey
}

func (p *BdjEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewBdjEnvService
}

func (p *BdjEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{p.Folder}
}

func (p *BdjEnvProvider) IsDefer() bool {
	return false
}

func (p *BdjEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	p.Folder = app.BaseFolder()
	return nil
}
