package middleware

import (
	"github.com/theone-daxia/bdj/framework/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 捕获 c.Next() 出现的 panic
		defer func() {
			if err := recover(); err != nil {
				ctx.ISetStatus(500).IJson(err)
			}
		}()

		ctx.Next()
	}
}
