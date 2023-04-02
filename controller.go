package main

import (
	"github.com/theone-daxia/bdj/framework"
	"time"
)

type req struct {
	Uri  string
	Path string
}

func UserAController(ctx *framework.Context) error {
	ret := req{ctx.GetRequest().RequestURI, ctx.GetRequest().URL.Path}
	ctx.Json(ret)
	return nil
}

func SubjectGetController(ctx *framework.Context) error {
	time.Sleep(10 * time.Second)
	ctx.Json("10秒正常结束")
	return nil
}

func TestController(ctx *framework.Context) error {
	ctx.Json("TestController")
	return nil
}
