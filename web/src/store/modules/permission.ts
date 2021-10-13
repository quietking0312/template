import {AppRouteRecordRaw} from "@/router/types";
import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import store from "@/store";
import path from "path";
import {asyncRouterMap, constantRouterMap} from "@/router";
import {deepClone} from "@/utils";

export interface PermissionState {
    routers: AppRouteRecordRaw[]
    addRouters: AppRouteRecordRaw[]
    isAddRouters: boolean
}

@Module({dynamic: true, namespaced: true, store, name: 'permission'})
class Permission extends VuexModule implements PermissionState {
    public routers = [] as any[]
    public addRouters = [] as any[]
    public isAddRouters = false

    @Mutation
    private SET_ROUTERS(routers: AppRouteRecordRaw[]): void {
        // 动态路由，404一定要放到最后面
        this.addRouters = routers.concat([{
            path: '/:path(.*)*',
            redirect: '/404',
            name: '404',
            meta: {
                hidden: true,
                breadcrumb: false
            }
        }])
        // 渲染菜单的所有路由
        this.routers = deepClone(constantRouterMap, ['component']).concat(routers)
    }
    @Mutation
    private SET_ISADDROUTERS(state: boolean): void {
        this.isAddRouters = state
    }

    @Action
    public GenerateRoutes(): Promise<unknown> {
        return new Promise<void>(resolve => {
            // 路由权限控制
            const routerMap: AppRouteRecordRaw[] = generateRoutes(deepClone(asyncRouterMap, ['component']))
            this.SET_ROUTERS(routerMap)
            resolve()
        })
    }
    @Action
    public SetIsAddRouters(state: boolean): void {
        this.SET_ISADDROUTERS(state)
    }
}

// 路由过滤，主要用于权限控制
function generateRoutes(routes: AppRouteRecordRaw[], basePath = '/'): AppRouteRecordRaw[] {
    const res: AppRouteRecordRaw[] = []

    for (const route of routes) {
        // skip some route
        if (route.meta && route.meta.hidden && !route.meta.showMainRoute) {
            continue
        }

        let data: any = null


        data = Object.assign({}, route)

        // recursive child routes
        if (route.children && data) {
            data.children = generateRoutes(route.children, path.resolve(basePath, data.path))
        }
        if (data) {
            res.push(data as AppRouteRecordRaw)
        }
    }
    return res
}

export const permissionStore = getModule<Permission>(Permission)
