import type {App} from "vue";
import {createI18n} from "vue-i18n";

import elementZHCNLocale from 'element-plus/lib/locale/lang/zh-cn'
import zh_cn from "./zh_cn";
import wsCache, {cacheKey} from "../cache";

const messages = {
    zh_cn: {
        el: elementZHCNLocale.el,
        ...zh_cn
    }
}

export function getLanguage(): string {
    const chooseLanguage = wsCache.get(cacheKey.lang)
    if (chooseLanguage) return chooseLanguage
    return 'zh_cn'
}

const i18n = createI18n({
    locale: getLanguage(),
    messages
})

export function setupI18n(app: App<Element>) {
    app.use(i18n)
}

export default i18n
