import zh from './module/zh'
import de from './module/de'
import fr from './module/fr'
import tw from './module/tw'
import en from './module/en'
import { createI18n } from 'vue-i18n';

const i18n =createI18n({
    legacy: false,
    locale: localStorage.getItem('lang') || 'zh',
    globalInjection: true,
    messages: {
        zh,
        tw,
        en,
        de,
        fr,
    },
    warnHtmlMessage: false,
})

export default i18n;
