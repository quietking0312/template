import router from "@/router";

import NProgress from 'nprogress';

import 'nprogress/nprogress.css'
import {permissionStore} from "@/store/modules/permission";
import {RouteRecordRaw} from "vue-router";

NProgress.configure({ showSpinner: false })

const whiteList: string[] = ['/login']
router.beforeEach((to, from, next) => {
    NProgress.start()
    if (permissionStore.isAddRouters) {
        next()
        return
    }
    permissionStore.GenerateRoutes().then(() => {
        permissionStore.addRouters.forEach((route: RouteRecordRaw) => {
            router.addRoute(route)
        })
        const redirectPath = (from.query.redirect || to.path) as string
        const redirect = decodeURIComponent(redirectPath)
        const nextData = to.path === redirect? { ...to, replace: true }: {path: redirect}
        permissionStore.SetIsAddRouters(true)
        next(nextData)
    })
})

router.afterEach((to) => {
    document.title = `${to.meta.title}`
    NProgress.done()
})
