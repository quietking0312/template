import type { App } from "vue";
import permission from "./permission"

export function setupDirective(app: App<Element>) {
    app.directive("permission", permission)
}
