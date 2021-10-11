import store from "@/store";
import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import wsCache, {cacheKey} from "@/cache";

export interface AppState {
    collapsed: boolean // 菜单栏是否缩放
    lang: string // 多语言
    layout: string // layout 布局
    showLogo: boolean // 是否显示logo
}

@Module({ dynamic: true, namespaced: true, store, name: 'app'})
class App extends VuexModule implements AppState {
    public collapsed = false
    public lang = wsCache.get(cacheKey.lang)? wsCache.get(cacheKey.lang): 'zh_cn'
    public layout = 'Classic'
    public showLogo = false

    @Mutation
    private SET_COLLAPSED(collapsed: boolean): void {
        this.collapsed = collapsed
    }

    @Action
    public SetCollapsed(collapsed: boolean): void {
        this.collapsed = collapsed
    }

    @Mutation
    private SET_LANG(lang: string):void {
        this.lang = lang
    }

    @Action
    public  SetLang(lang: string): void {
        this.SET_LANG(lang)
    }


}

export const appStore = getModule<App>(App)
