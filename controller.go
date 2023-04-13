package main

import (
	"github.com/theone-daxia/bdj/framework/gin"
	"time"
)

type req struct {
	Uri  string
	Path string
}

func UserAController(ctx *gin.Context) {
	ret := req{ctx.Request.RequestURI, ctx.Request.URL.Path}
	ctx.ISetOkStatus().IJson(ret)
}

func SubjectGetController(ctx *gin.Context) {
	time.Sleep(10 * time.Second)
	ctx.ISetOkStatus().IJson("10秒正常结束")
}

func TestController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("TestController")
}
