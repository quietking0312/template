import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import { appStore } from "@/store/modules/app";
import wsCache, {cacheKey} from "@/cache";
import router from "@/router";

const whiteList: string[] = ['/login'] // 不重定向白名单
router.beforeEach((to, from, next) => {
    NProgress.start()
    console.log(appStore.collapsed)
    if (wsCache.get(cacheKey.userInfo) || true) {
        if (to.path === '/login') {
            next({ path: '/'})
        } else {
            next()
            return
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
