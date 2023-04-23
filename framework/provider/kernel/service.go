package kernel

import (
	"github.com/theone-daxia/bdj/framework/gin"
	"net/http"
)

type BdjKernelService struct {
	engine *gin.Engine
}

func NewBdjKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return BdjKernelService{engine: httpEngine}, nil
}

func (s *BdjKernelService) HttpEngine() http.Handler {
	return s.engine
}
