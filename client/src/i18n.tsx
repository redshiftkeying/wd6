import LanguageDetector from 'i18next-browser-languagedetector';
import i18n from "i18next";
import enUsTrans from "./locales/en-us.json";
import zhCnTrans from "./locales/zh-cn.json";
import { initReactI18next } from 'react-i18next';

i18n.use(LanguageDetector) 
    .use(initReactI18next) //init i18next
    .init({
        //引入资源文件
        resources: {
            en: {
                translation: enUsTrans,
            },
            zh: {
                translation: zhCnTrans,
            },
        },
        fallbackLng: "en",
        debug: false,
        interpolation: {
            escapeValue: false, // not needed for react as it escapes by default
        },
    })

export default i18n;