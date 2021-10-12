import type { RouteRecordRaw } from "vue-router";

export interface RouteMeta {
    hidden?: boolean,
    title?: string,
    icon?: string,
    alwaysShow?: boolean,
}

export interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta'> {
    meta: RouteMeta
    children?: Array<AppRouteRecordRaw>
}

