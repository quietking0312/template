/**
 * Options API形式的组件生命周期钩子和Composition API之间的实际对应关系
 * beforeCreate -> setup()
 * created -> setup()
 * beforeMount -> onBeforeMount
 * mounted -> onMounted
 * beforeUpdate -> onBeforeUpdate
 * updated -> onUpdated
 * beforeDestroy -> onBeforeUnmount
 * destroyed -> onUnmounted
 * errorCaptured -> onErrorCaptured
 */

import '@/plugins/svgins'
import { createApp } from 'vue';
import App from './App.vue';
import {setupStore} from "./store";
import router, {setupRouter} from "./router";
import {setupI18n} from "./lang";
// import "element-plus/theme-chalk/index.css" // 部分模块使用了单独引入，ElementPlusResolver 插件无法检测，需要引入样式
import "@/styles/reset.css";
import "@/styles/index.less";
import {setupGlobCom} from "@/components";
import './permission'
import {setupDirective} from "@/directive";

const setupAll = async () => {
    const app = createApp(App)
    setupRouter(app)
    setupStore(app)
    setupI18n(app)
    setupGlobCom(app)
    setupDirective(app)
    router.isReady().then(() => {
        app.mount('#app')
    })
}
setupAll()


