package middleware

import (
	"github.com/gin-gonic/gin"
	"server/core/logic"
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
				info, err := logic.LoginLogicObj.GetLoginUserInfo(token)
				if err != nil {
					resp.JSON(c, resp.ErrTokenExpire, err.Error(), nil)
					c.Abort()
					return
				}
				for _, pid := range info.PermissionIds {
					if pid == define.AdminPid {
						c.Next()
						return
					} else if pid == r.PermissionId {
						c.Next()
						return
					}
				}
				resp.JSON(c, resp.ErrPermission, "", nil)
				return
			}
		}
		c.Next()
	}
}
