package middleware

import (
	"context"
	"fmt"
	"github.com/theone-daxia/bdj/framework/gin"
	"log"
	"time"
)

func Timeout(duration time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), duration)
		defer cancel()

		finishCh := make(chan struct{}, 1)   // 负责通知结束
		panicCh := make(chan interface{}, 1) // 负责通知 panic

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicCh <- p
				}
			}()
			ctx.Next()
			finishCh <- struct{}{}
		}()

		select {
		case p := <-panicCh:
			ctx.ISetStatus(500).IJson("time out")
			log.Println(p)
		case <-finishCh:
			fmt.Println("finish")
		case <-durationCtx.Done():
			//ctx.ISetHasTimeout()
			ctx.ISetStatus(500).IJson("time out")
		}
	}
}
