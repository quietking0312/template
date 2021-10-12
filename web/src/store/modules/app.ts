import store from "../index";
import {getModule, Module, VuexModule} from "vuex-module-decorators";

export interface AppState {
    collapsed: boolean
    layout: string
}

@Module({ dynamic: true, namespaced: true, store, name: 'app'})
class App extends VuexModule implements AppState {
    public collapsed = false //菜单栏是否缩放
    public layout = "Classic"
}

export const appStore = getModule<App>(App)
