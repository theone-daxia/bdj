package http

import (
	"github.com/theone-daxia/bdj/app/http/module/demo"
	"github.com/theone-daxia/bdj/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	_ = demo.Register(r)
}
