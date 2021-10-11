import {AppRouteRecordRaw} from "@/router/types";
import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import store from "@/store";
import path from 'path'
import {asyncRouteMap, constantRouterMap} from "@/router";
import {deepClone} from "@/utils";


export interface PermissionState {
    routers: AppRouteRecordRaw[]
    addRouters: AppRouteRecordRaw[]
    isAddRouters: boolean
}

@Module({dynamic: true, namespaced: true, store, name:'permission'})
class Permission extends VuexModule implements PermissionState {
    public routers = [] as any[]
    public addRouters = [] as any[]
    public isAddRouters = false

    @Mutation
    private SET_ROUTERS(routers: AppRouteRecordRaw[]):void {
        this.addRouters = routers.concat([{
            path: '/:path(.*)*',
            redirect: '/404',
            name: '404',
            meta: {
                hidden: true,
                breadcrumb: false
            }
        }])
        this.routers = deepClone(constantRouterMap, ['component']).concat(routers)
    }

    @Mutation
    private SET_ISADDROUTERS(state: boolean): void {
        this.isAddRouters = state
    }
    @Action
    public GenerateRoutes(): Promise<unknown> {
        return new Promise<void>(resolve => {
            const routerMap: AppRouteRecordRaw[] = generateRoutes(deepClone(asyncRouteMap, ['component']))
            this.SET_ROUTERS(routerMap)
            resolve()
        })
    }

    @Action
    public SetIsAddrouters(state: boolean): void {
        this.SET_ISADDROUTERS(state)
    }
}

function generateRoutes(routes: AppRouteRecordRaw[], basePath='/'): AppRouteRecordRaw[] {
    const res: AppRouteRecordRaw[] = []
    for (const route of routes) {
        if (route.meta && route.meta.hidden) {
            continue
        }

        let data: any = null

        data = Object.assign({}, route)

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
