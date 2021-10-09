import { createStore } from 'vuex'
import type { App } from 'vue'

const store = createStore({
  modules: {}
})

export function setupStore(app: App<Element>): void {
  app.use(store)
}

export default store
