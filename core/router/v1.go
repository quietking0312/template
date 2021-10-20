package router

import (
	"net/http"
	v1 "server/core/router/api/v1"
	"server/core/router/api/v1/helloworld"
	"server/core/router/api/v1/login"
	"server/core/utils/define"
)

var v1RouteGroup = define.RouteGroup{
	Path: "v1",
	Routes: []*define.Route{
		{Path: "helloworld", Method: http.MethodGet, Handler: helloworld.GetHelloWorld, Title: "helloWorld"},
		{Path: "appinfo", Method: http.MethodGet, Handler: v1.GetAppInfo, Title: "应用信息"},
	},
	RouteGroups: []*define.RouteGroup{
		{
			Path: "login",
			Routes: []*define.Route{
				{Path: "", Method: http.MethodPost, Handler: login.Login, Title: "login"},
			},
		},
	},
}
