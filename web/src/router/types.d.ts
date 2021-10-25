import type { RouteRecordRaw } from "vue-router";

/**
 *  meta : {
    hidden: true              当设置 true 的时候该路由不会再侧边栏出现 如404，login等页面(默认 false)
    alwaysShow: true          当你一个路由下面的 children 声明的路由大于1个时，自动会变成嵌套的模式，
                              只有一个时，会将那个子路由当做根路由显示在侧边栏，
                              若你想不管路由下面的 children 声明的个数都显示你的根路由，
                              你可以设置 alwaysShow: true，这样它就会忽略之前定义的规则，
                              一直显示根路由(默认 false)
    title: 'title'            设置该路由在侧边栏和面包屑中展示的名字
    icon: 'svg-name'          设置该路由的图标
    noCache: true             如果设置为true，则不会被 <keep-alive> 缓存(默认 false)
    breadcrumb: false         如果设置为false，则不会在breadcrumb面包屑中显示(默认 true)
    affix: true               如果设置为true，则会一直固定在tag项中(默认 false)
    noTagsView: true           如果设置为true，则不会出现在tag中(默认 false)
    activeMenu: '/dashboard'  显示高亮的路由路径
    followAuth: '/dashboard'  跟随哪个路由进行权限过滤
    showMainRoute: true       设置为true即使hidden为true，也依然可以进行路由跳转(默认 false)
    followRoute: '/dashboard' 为路由设置跟随其他路由的权限
    permission: [100000] 路由权限, 不设置为任意用户可访问
}
**/
export interface RouteMeta {
    hidden?: boolean,
    title?: string,
    icon?: string,
    alwaysShow?: boolean,
    noCache?: boolean,
    breadcrumb?: boolean,
    affix?: boolean
    activeMenu?: string,
    parent?: string,
    noTagsView?: boolean,
    followAuth?: string
    showMainRoute?: boolean,
    followRoute?: string,
    permission?: number[]
}

export interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta'> {
    meta: RouteMeta
    children?: Array<AppRouteRecordRaw>
}

