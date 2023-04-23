package main

import (
	"github.com/theone-daxia/bdj/app/console"
	"github.com/theone-daxia/bdj/app/http"
	demoService "github.com/theone-daxia/bdj/app/provider/demo"
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/provider/app"
	"github.com/theone-daxia/bdj/framework/provider/kernel"
)

func main() {
	container := framework.NewBdjContainer()
	container.Bind(&app.BdjAppProvider{})
	container.Bind(&demoService.DemoServiceProvider{})

	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.BdjKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)
}
