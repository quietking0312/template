package define

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Route struct {
	BasePath     string `json:"base_path"`
	Path         string `json:"path"`
	Method       string `json:"method"`
	Handler      gin.HandlerFunc
	PermissionId uint32 `json:"permission_id"`
	Title        string `json:"title"`
}

type RouteGroup struct {
	Path        string   `json:"path"`
	Routes      []*Route `json:"routes"`
	RouteGroups []*RouteGroup
}

// FilterRoute 过滤路由, 返回true 添加， 返回false 不添加， 用于不同环境的处理
type FilterRoute func(path string, method string) bool

func (g *RouteGroup) GinGroup(r *gin.RouterGroup, filter FilterRoute) {
	gg := r.Group(g.Path)
	for _, route := range g.Routes {
		route.GinRoute(gg, filter)
	}
	for _, group := range g.RouteGroups {
		group.GinGroup(gg, filter)
	}
}

func (r *Route) GinRoute(g *gin.RouterGroup, filter FilterRoute) {
	r.BasePath = g.BasePath()
	path := strings.Join([]string{strings.TrimRight(r.BasePath, "/"), strings.Trim(r.Path, "/")}, "/")
	if filter(path, r.Method) {
		g.Handle(r.Method, r.Path, func(c *gin.Context) {
			//defer func() {
			//	if panicValue := recover(); panicValue != nil {
			//		msg := ""
			//		fmt.Println(fmt.Errorf("%s %s: %v", r.Method, path, panicValue))
			//		for i := 1; ; i++ {
			//			pc, file, line, ok := runtime.Caller(i)
			//			if !ok {
			//				break
			//			}
			//			msg = fmt.Sprintf("%s %s:%d(0x%x)", msg, file, line, pc)
			//		}
			//		log.Error(fmt.Sprintf("%s %s", r.Method, path),
			//			zap.Error(fmt.Errorf("%v", panicValue)),
			//			zap.Error(fmt.Errorf("%s", msg)))
			//		resp.JSON(c, resp.ErrServer, fmt.Sprintf("%v", panicValue), "") // 确保服务错误 前期有返回
			//	}
			//}()
			r.Handler(c)
		})
	}
}
