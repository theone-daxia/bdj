package middleware

import "github.com/theone-daxia/bdj/framework"

func Recovery() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		// 捕获 c.Next() 出现的 panic
		defer func() {
			if err := recover(); err != nil {
				ctx.SetStatus(500).Json(err)
			}
		}()

		ctx.Next()

		return nil
	}
}
