import type { App } from "vue";
// import { createStore } from "vuex";
import { createPinia } from "pinia";

// const store = createStore({
//     modules: {}
// })
const store = createPinia()

export function setupStore(app: App<Element>) {
    app.use(store)
}

export { store }
