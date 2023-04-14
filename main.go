package main

import (
	"context"
	"github.com/theone-daxia/bdj/app/provider/demo"
	"github.com/theone-daxia/bdj/framework/gin"
	"github.com/theone-daxia/bdj/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	core.Bind(&demo.DemoServiceProvider{})
	core.Use(gin.Recovery(), middleware.Cost())
	registerRouter(core)
	server := http.Server{
		Addr:    ":8888",
		Handler: core,
	}
	go func() {
		_ = server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	// 优雅关停（不超过 5s）
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatalln("server shutdown:", err)
	}
}
