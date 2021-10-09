import { createApp } from 'vue'
import App from './App.vue'
import router, { setupRouter } from './router'
import {setupStore} from './store'
import './permission'
import { setupElement } from "@/libs/element";
import {setupI18n} from "@/lang";

const app = createApp(App)


setupRouter(app)

setupStore(app)

setupElement(app)

setupI18n(app)

router.isReady().then(() => {
    app.mount('#app')
})
