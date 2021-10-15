import {AppRouteRecordRaw} from "./types";
import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";
import type { App } from "vue";

const Layout = () => import('../layout/index.vue')

/**
 * redirect: noredirect        当设置 noredirect 的时候该路由在面包屑导航中不可被点击
 * name:'router-name'          设定路由的名字，一定要填写不然使用<keep-alive>时会出现各种问题
 * meta : {
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
  }
 **/


export const constantRouterMap: AppRouteRecordRaw[] = [
    {
      path: '/redirect',
      component: Layout,
      children: [
          {
              path: '/redirect/:path*',
              component:() => import('@/components/Redirect/index.vue'),
              meta: {}
          }
      ],
        meta: { hidden: true }
    },
    {
        path: '/login',
        component: () => import('@/views/login/index.vue'),
        name: 'Login',
        meta: {hidden: true, title: '登录', noTagsView: true}
    },
    {
        path: "/",
        component: Layout,
        redirect: '/dashboard',
        children:[{
            path: 'dashboard',
            name: "Dashboard",
            component: () => import('@/views/dashboard/index.vue'),
            meta: {title: "dashboard", icon: "dashboard", affix: true}

        }],
        meta: { }
    }
]

export const asyncRouterMap: AppRouteRecordRaw[] = [
    {
        path: "/demo",
        component: Layout,
        redirect: "/demo/menu1",
        children: [{
            path: 'menu1',
            name: 'menu1Demo',
            component: () => import('@/views/demo/Menu1.vue'),
            meta: { title: "menu1Demo" }
        },{
            path: 'menu2',
            name: 'menu2Demo',
            component: () => import('@/views/demo/Menu2.vue'),
            meta: { title: "menu2Demo" }
        }],
        meta: { title: "demo", icon: "nested" }
    }
]

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
