package middleware

import (
	"github.com/theone-daxia/bdj/framework/gin"
	"log"
	"time"
)

func Cost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		cost := time.Since(startTime)
		log.Printf("uri: %v, cost: %v", ctx.Request.RequestURI, cost)
	}
}
