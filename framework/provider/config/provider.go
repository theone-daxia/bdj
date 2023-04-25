package config

import (
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/contract"
	"path/filepath"
)

type BdjConfigProvider struct {
}

func (p *BdjConfigProvider) Name() string {
	return contract.ConfigKey
}

func (p *BdjConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewBdjConfig
}

func (p *BdjConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{c, envFolder, envService.All()}
}

func (p *BdjConfigProvider) IsDefer() bool {
	return false
}

func (p *BdjConfigProvider) Boot(c framework.Container) error {
	return nil
}
