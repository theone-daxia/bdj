package contract

import "net/http"

const KernelKey = "bdj:kernel"

// Kernel 接口提供框架最核心的结构
type Kernel interface {
	// HttpEngine http.Handler结构，作为net/http框架使用，实际上是gin.Engine
	HttpEngine() http.Handler
}
