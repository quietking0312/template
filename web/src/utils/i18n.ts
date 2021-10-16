import { useI18n } from "vue-i18n";

export function generateTitle(title: string): string {
    const { t, te } =  useI18n()
    const hasKey = te('route.' + title)
    if (hasKey) {
        return t('route.' + title)
    }
    return title
}
