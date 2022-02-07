import router from "@/router";

import NProgress from 'nprogress';

import 'nprogress/nprogress.css'
import {usePermissionStoreWithOut} from "@/store/modules/permission";
import {RouteRecordRaw} from "vue-router";
import wsCache, {cacheKey} from "@/cache";
import {useUserInfoStoreWithOut} from "@/store/modules/userInfo";

NProgress.configure({ showSpinner: false })
const permissionStore = usePermissionStoreWithOut()
const userInfoStore = useUserInfoStoreWithOut()
const whiteList: string[] = ['/login']
router.beforeEach(async (to, from, next) => {
    NProgress.start()
    if (wsCache.get(cacheKey.userInfo)) {
        if (to.path === '/login') {
            next({path: '/'})
        } else {
            if (permissionStore.getIsAddRouters) {
                next()
                return
            } else {
                try {
                    let permissionIdList = [] as number[]
                    await userInfoStore.SetUserInfo().then(permission_id => {
                        permissionIdList = permission_id as number[]
                    })
                    await permissionStore.GenerateRoutes(permissionIdList).then(() => {
                        permissionStore.getAddRouters.forEach(async (route) => {
                            await router.addRoute(route as RouteRecordRaw)
                        })
                        const redirectPath = (from.query.redirect || to.path) as string
                        const redirect = decodeURIComponent(redirectPath)
                        const nextData = to.path === redirect ? {...to, replace: true} : {path: redirect}
                        permissionStore.SetIsAddRouters(true)
                        next(nextData)
                    })
                }catch (err) {
                    await userInfoStore.resetToken()
                    next({path: '/login', query: {redirect: to.path}})
                    NProgress.done()
                }
            }
        }
    } else {
        if (whiteList.indexOf(to.path) !== -1) {
            next()
        } else {
            next({path: '/login', query: {redirect: to.path}})
        }
    }
})

router.afterEach((to) => {
    document.title = `${to.meta.title}`
    NProgress.done()
})
