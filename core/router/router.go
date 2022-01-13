package router

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os"
	"os/signal"
	"server/common/log"
	"server/core/config"
	"server/core/utils/define"
	"server/core/utils/middleware"
	"strings"
	"syscall"
)

// NewHTTPRouter 创建HTTP路由
func NewHTTPRouter() *http.Server {
	authMiddle := middleware.NewAuthMiddleWare()
	var router *gin.Engine
	if config.GetConfig().Server.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
		router = gin.New()
		router.Use(middleware.Cors())
		router.Use(authMiddle.Auth())
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
		router.Use(authMiddle.Auth())
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
		if config.GetConfig().Server.Mode != "debug" && strings.Index(path, "helloworld") < 0 {
			return false
		}
		return true
	})
	// 由于路径拼接问题，路由map需要执行完 GinGroup 才会生成
	define.DefaultPermissionList.SetData(v1RouteGroup)
	authMiddle.SetPermission(define.DefaultPermissionList.RouteMap(v1RouteGroup))
	httpSev := &http.Server{
		Handler: router,
	}
	return httpSev
}

func GetListen(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}

// Run 使用cmux多路复用启动器
func Run(lis net.Listener) error {

	m := cmux.New(lis)
	httpServer := NewHTTPRouter()
	httpL := m.Match(cmux.HTTP1Fast())
	go func() {
		_ = httpServer.Serve(httpL)
	}()
	log.Info("start run http server", zap.String("addr", lis.Addr().String()))
	if err := m.Serve(); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return m.Serve()
}

// RunHttpServer 使用http包启动,支持优雅退出
func RunHttpServer(lis net.Listener) error {
	httpServer := NewHTTPRouter()

	go func() {
		if err := httpServer.Serve(lis); err != nil && err != http.ErrServerClosed {
			log.Fatal("", zap.Error(err))
		}
	}()
	log.Info("start run http server", zap.String("addr", lis.Addr().String()))
	shutdown(httpServer)
	return nil
}

func shutdown(httpServer *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-quit:
		fmt.Println(sig)
		if err := httpServer.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
		}
	}
}
