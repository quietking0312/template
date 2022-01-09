package define

import (
	"fmt"
	"strings"
)

var DefaultPermissionList = NewPermission()

type RouteItem struct {
	Path         string `json:"path"`
	Method       string `json:"method"`
	PermissionId uint32 `json:"permission_id"`
	Title        string `json:"title"`
}

type PermissionList struct {
	data         []RouteItem
	pidIndexData map[uint32]RouteItem
}

func NewPermission() *PermissionList {
	return &PermissionList{
		data:         []RouteItem{},
		pidIndexData: make(map[uint32]RouteItem),
	}
}

// SetData 写入路由列表
func (p *PermissionList) SetData(g RouteGroup) {
	for _, route := range g.Routes {
		path := strings.Join([]string{strings.TrimRight(route.BasePath, "/"), strings.Trim(route.Path, "/")}, "/")
		var r = RouteItem{
			Path:         strings.TrimRight(path, "/"),
			Method:       route.Method,
			PermissionId: route.PermissionId,
			Title:        route.Title,
		}
		if r.PermissionId != 0 {
			p.data = append(p.data, r)
			if _, o := p.pidIndexData[r.PermissionId]; !o {
				p.pidIndexData[r.PermissionId] = r
			} else {
				panic(fmt.Sprintf("pid:%d is exists", r.PermissionId))
			}
		}
	}
	for _, group := range g.RouteGroups {
		p.SetData(*group)
	}

	return
}

// RouteMap 获取路由map
// return map[path][method]RoutePermission
func (p *PermissionList) RouteMap(g RouteGroup) map[string]map[string]RouteItem {
	var permissionMap = make(map[string]map[string]RouteItem)
	if len(p.data) == 0 {
		p.SetData(g)
	}
	for _, route := range p.data {
		pM, ok := permissionMap[route.Path]
		if !ok {
			pM = make(map[string]RouteItem)
			permissionMap[route.Path] = pM
		}
		pM[route.Method] = route
	}
	return permissionMap
}

func (p PermissionList) GetList() []RouteItem {
	return p.data
}

func (p PermissionList) PidIsExists(pid uint32) bool {
	_, ok := p.pidIndexData[pid]
	return ok
}
