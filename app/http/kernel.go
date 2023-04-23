package http

import (
	"github.com/theone-daxia/bdj/framework/gin"
)

// NewHttpEngine 创建一个绑定了路由的 web 引擎
func NewHttpEngine() (*gin.Engine, error) {
	// 设置为 release，即默认启动不输出调试信息
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 绑定路由
	Routes(r)
	return r, nil
}
