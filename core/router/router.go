package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
	"server/core/utils/middleware"
)

// NewHTTPRouter 创建HTTP路由
func NewHTTPRouter() *http.Server {
	router := gin.New()
	router.Use(middleware.Cors())
	router.Use(gin.Logger())
	apiGroup := router.Group("api")
	// 添加服务信息监听
	pprof.RouteRegister(apiGroup)

	v1RouteGroup.GinGroup(apiGroup, func(path string, method string) bool {
		return true
	})

	httpSev := &http.Server{
		Handler: router,
	}
	return httpSev
}

func GetListen(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}

func Run(lis net.Listener) error {

	m := cmux.New(lis)
	httpServer := NewHTTPRouter()
	httpL := m.Match(cmux.HTTP1Fast())
	go func() {
		_ = httpServer.Serve(httpL)
	}()

	return m.Serve()
}
