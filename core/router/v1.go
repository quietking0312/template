package router

import (
	"net/http"
	v1 "server/core/router/api/v1"
	"server/core/utils/define"
)

var v1RouteGroup = define.RouteGroup{
	Path: "v1",
	Routes: []*define.Route{
		{Path: "helloworld", Method: http.MethodGet, Handler: v1.GetHelloWorld, Title: "helloWorld"},
		{Path: "appinfo", Method: http.MethodGet, Handler: v1.GetAppInfo, Title: "应用信息"},
	},
	RouteGroups: []*define.RouteGroup{
		{
			Path: "login",
			Routes: []*define.Route{
				{Path: "", Method: http.MethodPost, Handler: v1.Login, Title: "login"},
			},
		}, {
			Path: "userinfo",
			Routes: []*define.Route{
				{Path: "", Method: http.MethodGet, Handler: v1.GetUserInfo, Title: "用户信息"},
			},
		}, {
			Path: "permission",
			Routes: []*define.Route{
				{Path: "user/list", Method: http.MethodGet, Handler: v1.GetUserListApi, Title: "获取用户列表"},
				{Path: "user", Method: http.MethodPost, Handler: v1.PostUserApi, Title: "添加用户"},
				{Path: "user", Method: http.MethodPut, Handler: v1.PutUserApi, Title: "修改用户"},
				{Path: "user", Method: http.MethodDelete, Handler: v1.DeleteUserApi, Title: "删除用户"},
				{Path: "role/list", Method: http.MethodGet, Handler: v1.GetRoleListApi, Title: "获取角色列表"},
				{Path: "role", Method: http.MethodPost, Handler: v1.PostRoleApi, Title: "添加角色"},
				{Path: "role", Method: http.MethodPut, Handler: v1.PutRoleApi, Title: "修改角色"},
				{Path: "role", Method: http.MethodDelete, Handler: v1.DeleteRoleApi, Title: "删除角色"},
			},
		},
	},
}
