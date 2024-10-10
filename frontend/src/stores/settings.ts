import { defineStore } from 'pinia'
import { locales } from '@/locales'
import { get, set, split, cloneDeep } from 'lodash'
import { i18nGlobal } from '@/utils/i18n.js'
import { h, nextTick } from 'vue'
import { enUS, NButton, NSpace, useOsTheme, zhCN } from 'naive-ui'
import * as FileStorage from '@/../wailsjs/go/main/ConfigFileStorage'

export const useSettingsStore = defineStore('settings', {
    state: () => ({
        general: {
            language: "auto",
            theme: "axolotl-light",
        },
        editor: {},
        previousSettingsState: {},
    }),
    getters: {
        optionsTheme() {
            return [
                {
                    value: 'axolotl-light',
                    label: 'settings.general.axolotl_light',
                },
                {
                    value: 'scarecrow-dark',
                    label: 'settings.general.scarecrow_dark',
                },
            ]
        },
        optionsLocales() {
            const options = Object.entries(locales).map(([key, value]) => ({
                value: key,
                label: value['name'],
            }))
            options.splice(0, 0, {
                value: 'auto',
                label: 'settings.general.system_lang',
            })
            return options
        },
        /**
         * Get current language shortcode, e.g. 'en'
         * If used 'auto' - we try to detect language from a navigator
         * @return {string}
         */
        currentLanguageCode(): string {
            let lang = get(this.general, 'language', 'auto')
            if (lang === 'auto' && navigator.language) {
                lang = split(navigator.language, '-')[0]
            }
            return lang || 'en'
        },
        themeLocale() {
            const lang = this.currentLanguageCode
            switch (lang) {
                case 'zh':
                    return zhCN
                default:
                    return enUS
            }
        },
    },
    actions: {
        _setPreferences(data: any) {
            for (const key in data) {
                set(this, key, data[key])
            }
        },

        async loadPreferences(): Promise<void> {
            const { success, data } = await FileStorage.Get("settings.json", "null");

            console.log(success);
            console.log(data);
            if (success) {
                this.previousSettingsState = cloneDeep(data)
                // this._setPreferences(data)
                // TODO: set default values here
            
                // i18nGlobal.locale.value = this.currentLanguageCode
            }
        },
    }
})