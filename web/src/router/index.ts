import {AppRouteRecordRaw} from "./types";
import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";
import type { App } from "vue";

export const constantRouterMap: AppRouteRecordRaw[] = [
    {
        path: "/",
        component: () => import('@/views/dashboard/index.vue'),
        name: "Dashboard",
        meta: { title: "dashboard", icon: "dashboard" }
    }
]

export const asyncRouterMap: AppRouteRecordRaw[] = []

const router = createRouter({
    history: createWebHistory(),
    strict: true,
    routes: constantRouterMap as RouteRecordRaw[]
})

export function resetRouter(): void{
    const resetWhiteNameList = [
        'Login'
    ]
    router.getRoutes().forEach((route) => {
        const { name } = route
        if (name && !resetWhiteNameList.includes(name as string)) {
            router.hasRoute(name) && router.removeRoute(name)
        }
    })
}

export function setupRouter(app: App<Element>) {
    app.use(router)
}
export default router
