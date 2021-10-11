import type { App } from "vue";
import SvgIcon from './SvgIcon/index.vue'

import '@/assets/icons'

export function setupGlobCom(app: App<Element>): void {
    app.component('SvgIcon', SvgIcon)
}
