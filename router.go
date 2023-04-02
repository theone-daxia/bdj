package main

import (
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/middleware"
)

func registerRouter(core *framework.Core) {

	// HTTP方法 + 静态路由
	core.Get("/user/a", UserAController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	//subjectApi.Use(middleware.Test1())
	{
		// 动态路由
		subjectApi.Get("/:id", middleware.Test2(), SubjectGetController)

		subGroup1 := subjectApi.Group("/a")
		subGroup2 := subGroup1.Group("/b")
		subGroup2.Get("/test", TestController)
	}

}
