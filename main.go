package main

import (
	"github.com/theone-daxia/bdj/app/console"
	"github.com/theone-daxia/bdj/app/http"
	demoService "github.com/theone-daxia/bdj/app/provider/demo"
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/provider/app"
	"github.com/theone-daxia/bdj/framework/provider/config"
	"github.com/theone-daxia/bdj/framework/provider/env"
	"github.com/theone-daxia/bdj/framework/provider/kernel"
)

func main() {
	container := framework.NewBdjContainer()
	container.Bind(&app.BdjAppProvider{})
	container.Bind(&demoService.DemoServiceProvider{})
	container.Bind(&env.BdjEnvProvider{})
	container.Bind(&config.BdjConfigProvider{})

	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.BdjKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)
}
