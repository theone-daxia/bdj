package command

import (
	"context"
	"github.com/theone-daxia/bdj/framework/cobra"
	"github.com/theone-daxia/bdj/framework/contract"
	"github.com/theone-daxia/bdj/framework/provider/kernel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// appCmd 是命令行参数第一级为 app 的命令，它没有实际功能，只是打印帮助文档
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 打印帮助文档
		cmd.Help()
		return nil
	},
}

// appStartCmd 启动一个 web 服务
var appStartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动一个 web 服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 从 command 中获取服务容器
		container := cmd.GetContainer()
		// 从容器中获取 kernel 服务实例
		kernelService := container.MustMake(contract.KernelKey).(kernel.BdjKernelService)
		// 从 kernel 服务实例中获取引擎
		engine := kernelService.HttpEngine()

		server := &http.Server{
			Handler: engine,
			Addr:    ":8888",
		}

		// 启动服务的 goroutine
		go func() {
			_ = server.ListenAndServe()
		}()

		// 当前 goroutine 等待的信号量
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit

		// 优雅关停
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}

		return nil
	},
}

// initAppCommand 初始化 app 命令和其子命令
func initAppCommand() *cobra.Command {
	appCmd.AddCommand(appStartCmd)
	return appCmd
}
