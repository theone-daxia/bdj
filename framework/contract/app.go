package contract

// AppKey 定义字符串凭证
const AppKey = "bdj:app"

type App interface {
	// Version 当前版本
	Version() string

	// BaseFolder 项目根目录
	BaseFolder() string

	// ConfigFolder 配置文件目录
	ConfigFolder() string

	// LogFolder 日志文件目录
	LogFolder() string

	// ProviderFolder 业务自己的服务提供者目录
	ProviderFolder() string

	// MiddlewareFolder 业务自己的中间件目录
	MiddlewareFolder() string

	// CommandFolder 业务自定义的命令目录
	CommandFolder() string

	// RuntimeFolder 业务运行中间态信息目录
	RuntimeFolder() string

	// TestFolder 存放测试用例、测试数据的目录
	TestFolder() string
}
