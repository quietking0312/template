import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import type { App } from 'vue'
import Home from '../views/Home.vue'
import {AppRouteRecordRaw} from "@/router/types";


const Layout = () => import('../layout/index.vue')

export const constantRouterMap: Array<AppRouteRecordRaw> = [
  {
    path: "/redirect",
    component: Layout,
    meta: {},
    children: [
      {
        path: '/redirect/:path*',
        component: () => import('@c/Redirect/index.vue'),
        meta: {}
      }
    ]
  },
  {
    path: '/',
    name: 'Home',
    component: Layout,
    redirect: '/home',
    meta: {},
    children: [
      {
        path: 'home',
        name: 'Home',
        component: Home,
        meta: {}
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: {}
  }
]

export const asyncRouteMap: Array<AppRouteRecordRaw> = [

]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes: constantRouterMap as RouteRecordRaw[]
})

export function resetRouter():void {
  const resetWhiteNameList = [
      'Login'
  ]
  router.getRoutes().forEach((route) => {
    const  { name } = route
    if (name && !resetWhiteNameList.includes(name as string)) {
      router.hasRoute(name) && router.removeRoute(name)
    }
  })
}

export function setupRouter(app: App<Element>): void {
  app.use(router)
}

export default router
