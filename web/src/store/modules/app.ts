import { store } from "@/store";
import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import {getLanguage} from "@/lang";
import wsCache, {cacheKey} from "@/cache";
import {defineStore} from "pinia";

export type LayoutType = 'Classic' | 'LeftTop' | 'Top'

export interface AppState {
    collapsed: boolean
    showTags: boolean
    showLogo: boolean
    showNavbar: boolean
    fixedHeader: boolean
    layout: LayoutType
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

export const useAppStore = defineStore({
    id: "app",
    state: (): AppState => ({
        collapsed: false, //菜单栏是否缩放
        showTags: false, // 是否显示标签栏
        showLogo: true, // 是否显示logo
        showNavbar: true, // 是否显示navbar
        fixedHeader: true, // 是否固定header
        layout: "Classic", // layout布局
        showBreadcrumb: false, // 是否显示面包屑
        showHamburger: true, // 是否显示侧边栏缩收按钮
        showScreenfull: false, // 是否全屏按钮
        showUserInfo: true, // 是否显示用户头像
        title: "vue-admin", // 标题
        logoTitle: "vue-admin", // logo标题
        greyMode: false,  // 是否开始灰色模式，用于特殊悼念日
        showBackTop: true, // 是否显示回到顶部
        lang: getLanguage(),
        showLanguage: false, // 是否显示选择语言
    }),
    getters: {
        getCollapsed(): boolean {
            return this.collapsed
        },
        getShowLogo(): boolean {
            return this.showLogo
        },
        getShowTags(): boolean {
            return this.getShowTags
        },
        getShowNavbar(): boolean {
            return this.showNavbar
        },
        getFixedHeader(): boolean {
            return this.fixedHeader
        },
        getShowBreadcrumb(): boolean {
            return this.showBreadcrumb
        },
        getShowHamburger(): boolean {
            return this.showHamburger
        },
        getShowScreenfull(): boolean {
            return this.showScreenfull
        },
        getShowUserInfo(): boolean {
            return this.showUserInfo
        },
        getTitle(): string {
            return this.title
        },
        getLogoTitle(): string {
            return this.logoTitle
        },
        getGreyMode(): boolean {
            return this.getGreyMode
        },
        getLayout(): LayoutType {
            return this.layout
        },
        getShowBackTop(): boolean {
            return this.showBackTop
        },
        getShowLanguage(): boolean {
            return this.showLanguage
        },
        getLang(): string {
            return this.lang
        }
    },
    actions: {
        SetCollapsed(collapsed: boolean): void {
            this.collapsed =collapsed
        },
        SetLanguage(lang: string): void {
            wsCache.set(cacheKey.lang, lang)
            this.lang = lang
        }
    }
})

// @Module({ dynamic: true, namespaced: true, store, name: 'app'})
// class App extends VuexModule implements AppState {
//     public collapsed = false //菜单栏是否缩放
//     public showTags = false // 是否显示标签栏
//     public showLogo = true // 是否显示logo
//     public showNavbar = true // 是否显示navbar
//     public fixedHeader = true // 是否固定header
//     public layout = "Classic" // layout布局
//     public showBreadcrumb = false // 是否显示面包屑
//     public showHamburger = true // 是否显示侧边栏缩收按钮
//     public showScreenfull = false // 是否全屏按钮
//     public showUserInfo = true // 是否显示用户头像
//     public title = "vue-admin" // 标题
//     public logoTitle = "vue-admin" // logo标题
//     public greyMode = false  // 是否开始灰色模式，用于特殊悼念日
//     public showBackTop = true // 是否显示回到顶部
//     public lang = getLanguage()
//     public showLanguage = false // 是否显示选择语言
//
//     @Mutation
//     private SET_COLLAPSED(collapsed: boolean): void {
//         this.collapsed = collapsed
//     }
//
//     @Mutation
//     private SET_LANG(lang: string): void {
//         this.lang = lang
//     }
//
//     @Action
//     public SetCollapsed(collapsed: boolean): void {
//         this.SET_COLLAPSED(collapsed)
//     }
//
//     @Action
//     public SetLanguage(lang: string): void {
//         wsCache.set(cacheKey.lang, lang)
//         this.SET_LANG(lang)
//     }
// }
//
// export const appStore = getModule<App>(App)
export function useAppStoreWithOut() {
    return useAppStore(store)
}
