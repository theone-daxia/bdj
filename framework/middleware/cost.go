package middleware

import (
	"github.com/theone-daxia/bdj/framework"
	"log"
	"time"
)

func Cost() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		startTime := time.Now()
		ctx.Next()
		cost := time.Since(startTime)
		log.Printf("uri: %v, cost: %v", ctx.GetRequest().RequestURI, cost)
		return nil
	}
}
