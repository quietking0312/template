import type {App} from "vue";

import ElementConfig from './element.config'
import ElLoading, {ElAlert, ElAside, ElInfiniteScroll, ElMessageBox} from "element-plus";
import ElNotification from "element-plus";

const { size, zIndex } = ElementConfig

const components = [
    ElAlert,
    ElAside
]

const plugins = [
    ElInfiniteScroll,
    ElLoading,
    ElMessageBox,
    ElNotification
]

export function setupElement(app: App<Element>): void {
    components.forEach((component: any) => {
        app.component(component.name, component)
    })

    plugins.forEach((plugin: any) => {
        app.use(plugin)
    })

    app.config.globalProperties.$ELEMENT = { size: size, zIndex }
}
