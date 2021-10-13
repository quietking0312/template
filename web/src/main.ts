import { createApp } from 'vue';
import App from './App.vue';
import {setupStore} from "./store";
import router, {setupRouter} from "./router";
import {setupI18n} from "./lang";
import "@/styles/reset.css";
import "@/styles/index.less";
import {setupGlobCom} from "@/components";
import './permission'

const app = createApp(App)
setupRouter(app)
setupStore(app)
setupI18n(app)
setupGlobCom(app)

router.isReady().then(() => {
    app.mount('#app')
})

