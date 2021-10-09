import type {App} from "vue";
import { createI18n } from "vue-i18n";
import wsCache, {cacheKey} from "@/cache";
import elementZHLocale from 'element-plus/lib/locale/lang/zh-cn'
import zh_cnLocale from "@/lang/zh_cn";

const messages = {
    zh_cn: {
        el: elementZHLocale.el,
        ...zh_cnLocale
    }
}

export function getLanguage(): string {
    const lang = wsCache.get(cacheKey.lang)
    if (lang) return lang
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
