package router

import (
	"net/http"
	"server/core/router/api/v1/helloworld"
	"server/core/router/api/v1/login"
	"server/core/utils/define"
)

var v1RouteGroup = define.RouteGroup{
	Path: "v1",
	Routes: []*define.Route{
		{Path: "helloworld", Method: http.MethodGet, Handler: helloworld.GetHelloWorld, Title: "helloWorld"},
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
