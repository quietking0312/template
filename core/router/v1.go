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
		{Path: "notice", Method: http.MethodGet, Handler: v1.Notice, Title: "公告"},
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
				{Path: "user/list", Method: http.MethodGet, Handler: v1.GetUserListApi, Title: "获取用户列表", PermissionId: 100001},
				{Path: "user", Method: http.MethodPost, Handler: v1.PostUserApi, Title: "添加用户", PermissionId: 100002},
				{Path: "user", Method: http.MethodPut, Handler: v1.PutUserApi, Title: "修改用户", PermissionId: 100003},
				{Path: "user", Method: http.MethodDelete, Handler: v1.DeleteUserApi, Title: "删除用户", PermissionId: 100004},
				{Path: "role/list", Method: http.MethodGet, Handler: v1.GetRoleListApi, Title: "获取角色列表", PermissionId: 101001},
				{Path: "role", Method: http.MethodPost, Handler: v1.PostRoleApi, Title: "添加角色", PermissionId: 101002},
				{Path: "role", Method: http.MethodPut, Handler: v1.PutRoleApi, Title: "修改角色", PermissionId: 101003},
				{Path: "role", Method: http.MethodDelete, Handler: v1.DeleteRoleApi, Title: "删除角色", PermissionId: 101004},
				{Path: "", Method: http.MethodGet, Handler: v1.GetPermissionList, Title: "获取权限列表", PermissionId: 102001},
				{Path: "user/permission", Method: http.MethodPost, Handler: v1.PostUserPermission, Title: "修改用户权限", PermissionId: 102002},
				{Path: "role/permission", Method: http.MethodPost, Handler: v1.PostRolePermission, Title: "修改角色权限", PermissionId: 102003},
			},
		},
	},
}
