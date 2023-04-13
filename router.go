package main

import (
	"github.com/theone-daxia/bdj/framework/gin"
)

func registerRouter(core *gin.Engine) {

	// HTTP方法 + 静态路由
	core.GET("/user/a", UserAController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	//subjectApi.Use(middleware.Test1())
	{
		// 动态路由
		subjectApi.GET("/:id", SubjectGetController)

		subGroup1 := subjectApi.Group("/a")
		subGroup2 := subGroup1.Group("/b")
		subGroup2.GET("/test", TestController)
	}

}
