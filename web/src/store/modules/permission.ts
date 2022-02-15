import {AppRouteRecordRaw} from "@/router/types";
import { store } from "@/store";
import {asyncRouterMap, constantRouterMap} from "@/router";
import {deepClone} from "@/utils";
import {defineStore} from "pinia";
// @ts-ignore
import path from "path-browserify";
import permission from "@/directive/permission";

export interface PermissionState {
    routers: AppRouteRecordRaw[]
    addRouters: AppRouteRecordRaw[]
    isAddRouters: boolean
}

export const usePermissionStore = defineStore({
    id: 'permission',
    state: (): PermissionState => ({
        routers: [],
        addRouters: [],
        isAddRouters: false,
    }),
    getters: {
        getRouters(): AppRouteRecordRaw[] {
            return this.routers
        },
        getAddRouters(): AppRouteRecordRaw[] {
            return this.addRouters
        },
        getIsAddRouters(): boolean {
            return this.isAddRouters
        },
    },
    actions: {
        GenerateRoutes(permissionIdList: number[] = []): Promise<unknown> {
            return new Promise<void>(resolve => {
                // 路由权限控制
                const routerMap: AppRouteRecordRaw[] = generateRoutes(deepClone(asyncRouterMap, ['component']), permissionIdList)
                // 动态路由，404一定要放到最后面
                this.addRouters = routerMap.concat([{
                    path: '/:path(.*)*',
                    redirect: '/404',
                    name: '404',
                    meta: {
                        hidden: true,
                        breadcrumb: false
                    }
                }])
                // 渲染菜单的所有路由
                this.routers = deepClone(constantRouterMap, ['component']).concat(routerMap)
                resolve()
            })
        },
        SetIsAddRouters(state: boolean): void {
            this.isAddRouters = state
        }
    }
})

export function usePermissionStoreWithOut() {
    return usePermissionStore(store)
}

function hasPermission(permissionIdList: number[], route: AppRouteRecordRaw):boolean{
    if (route.meta && route.meta.permission) {
        return permissionIdList.some(permissionId => {
            if (permissionId === 100000) {
                return true
            } else {
                return route.meta.permission?.includes(permissionId)
            }
        })
    } else {
        return true
    }
}

// 路由过滤，主要用于权限控制
function generateRoutes(routes: AppRouteRecordRaw[], permissionIdList: number[] = [], basePath = '/'): AppRouteRecordRaw[] {

    const res: AppRouteRecordRaw[] = []

    for (const route of routes) {
        // skip some route
        if (route.meta && route.meta.hidden && !route.meta.showMainRoute) {
            continue
        }
        if (!hasPermission(permissionIdList, route)){
            continue
        }

        let data: any = null


        data = Object.assign({}, route)

        // recursive child routes
        if (route.children && data) {
            data.children = generateRoutes(route.children, permissionIdList, path.resolve(basePath, data.path))
        }
        if (data) {
            res.push(data as AppRouteRecordRaw)
        }
    }
    return res
}
