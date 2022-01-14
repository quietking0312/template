import { useI18n } from "vue-i18n";

export function generateTitle(title: any): any {
    const { t, te } =  useI18n()
    const hasKey = te('route.' + title)
    if (hasKey) {
        return t('route.' + title)
    }
    return title
}
