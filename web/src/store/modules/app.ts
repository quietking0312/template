import store from "../index";
import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import {getLanguage} from "@/lang";
import wsCache, {cacheKey} from "@/cache";

export interface AppState {
    collapsed: boolean
    showTags: boolean
    showLogo: boolean
    showNavbar: boolean
    fixedHeader: boolean
    layout: string
    showBreadcrumb: boolean
    showHamburger: boolean
    showScreenfull: boolean
    showUserInfo: boolean
    title: string
    logoTitle: string
    greyMode: boolean
    showBackTop: boolean
    lang: string
    showLanguage: boolean
}

@Module({ dynamic: true, namespaced: true, store, name: 'app'})
class App extends VuexModule implements AppState {
    public collapsed = false //菜单栏是否缩放
    public showTags = true // 是否显示标签栏
    public showLogo = true // 是否显示logo
    public showNavbar = true // 是否显示navbar
    public fixedHeader = true // 是否固定header
    public layout = "Classic" // layout布局
    public showBreadcrumb = true // 是否显示面包屑
    public showHamburger = true // 是否显示侧边栏缩收按钮
    public showScreenfull = true // 是否全屏按钮
    public showUserInfo = true // 是否显示用户头像
    public title = "vue-admin" // 标题
    public logoTitle = "vue-admin" // logo标题
    public greyMode = false  // 是否开始灰色模式，用于特殊悼念日
    public showBackTop = true // 是否显示回到顶部
    public lang = getLanguage()
    public showLanguage = true // 是否显示选择语言

    @Mutation
    private SET_COLLAPSED(collapsed: boolean): void {
        this.collapsed = collapsed
    }

    @Mutation
    private SET_LANG(lang: string): void {
        this.lang = lang
    }

    @Action
    public SetCollapsed(collapsed: boolean): void {
        this.SET_COLLAPSED(collapsed)
    }

    @Action
    public SetLanguage(lang: string): void {
        wsCache.set(cacheKey.lang, lang)
        this.SET_LANG(lang)
    }
}

export const appStore = getModule<App>(App)
