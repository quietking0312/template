import type { RouteRecordRaw } from "vue-router";

export interface RouteMeta {
    title?: string // 名称
    icon?: string // icon
}

export interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta'> {
    meta: RouteMeta
    children?: AppRouteRecordRaw[]
}
