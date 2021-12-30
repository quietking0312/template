package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"server/common/log"
	"server/core/config"
	"server/core/utils/middleware"
)

// NewHTTPRouter 创建HTTP路由
func NewHTTPRouter() *http.Server {
	var router *gin.Engine
	if config.GetConfig().Server.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
		router = gin.New()
		router.Use(middleware.Cors())
		if config.GetConfig().Log.RouteLog {
			router.Use(gin.Logger())
		}
		// 添加服务信息监听
		pprof.Register(router)
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		router.Use(middleware.Cors())
		router.Use(middleware.Recover())
		if config.GetConfig().Log.RouteLog {
			router.Use(gin.Logger())
		}
		if config.GetConfig().Server.PPROF {
			// 添加服务信息监听
			pprof.Register(router)
		}
	}

	apiGroup := router.Group("api")

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
	log.Info("start run http server", zap.String("addr", lis.Addr().String()))
	return m.Serve()
}
