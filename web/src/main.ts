import { createApp } from 'vue'
import App from './App.vue'
import router, { setupRouter } from './router'
import {setupStore} from './store'
import './permission'
import { setupElement } from "@/libs/element";
import {setupI18n} from "@/lang";
import {setupGlobCom} from "@/components";



import '@/styles/reset.css'

import '@/styles/index.less'

const app = createApp(App)


setupRouter(app)

setupStore(app)

setupElement(app)

setupI18n(app)

setupGlobCom(app)

router.isReady().then(() => {
    app.mount('#app')
})
