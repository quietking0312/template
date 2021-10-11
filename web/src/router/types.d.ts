import type { RouteRecordRaw } from "vue-router";

export interface RouteMeta {
    title?: string // 名称
    icon?: string // icon
    hidden?: boolean  // 当设置 true 的时候该路由不会再侧边栏出现 如404，login等页面(默认 false)
    breadcrumb?: boolean // 如果设置为false，则不会在breadcrumb面包屑中显示(默认 true)
}

export interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta'> {
    meta: RouteMeta
    children?: AppRouteRecordRaw[]
}
