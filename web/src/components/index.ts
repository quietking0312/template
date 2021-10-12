import type { App } from 'vue'
import SvgIcon from './SvgIcon/index.vue'// svg组件

// import '@/assets/icons' // 引入svg图标
import 'vite-plugin-svg-icons/register';

export function setupGlobCom(app: App<Element>): void {
    app.component('SvgIcon', SvgIcon)
}
