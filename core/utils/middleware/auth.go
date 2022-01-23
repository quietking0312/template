package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core/utils/define"
	"server/core/utils/resp"
)

type authMiddleWare struct {
	permissionMap map[string]map[string]define.RouteItem
}

func NewAuthMiddleWare() *authMiddleWare {
	return &authMiddleWare{
		permissionMap: make(map[string]map[string]define.RouteItem),
	}
}

func (a *authMiddleWare) SetPermission(permissionMap map[string]map[string]define.RouteItem) {
	a.permissionMap = permissionMap
}

func (a *authMiddleWare) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if r, ok := a.permissionMap[c.Request.URL.Path][c.Request.Method]; ok {
			if r.PermissionId != 0 {
				token := define.GetToken(c)
				if token == "" {
					resp.JSON(c, resp.ErrTokenExpire, "", nil)
					c.Abort()
					return
				}
				fmt.Println(token)

			}
		}
		c.Next()
	}
}
