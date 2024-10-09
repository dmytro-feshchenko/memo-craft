import { createPinia } from 'pinia'
import { createApp, nextTick } from 'vue'
import App from './App.vue'
import './style.css';
import './configs/themes/axolotl-light.css'; // TODO: Replace with dynamic theming
import dayjs from 'dayjs'
import duration from 'dayjs/plugin/duration'
import relativeTime from 'dayjs/plugin/relativeTime'
import { i18n } from '@/utils/i18n.js'
import { useSettingsStore } from '@/stores/settings'

dayjs.extend(duration)
dayjs.extend(relativeTime)

// import { bind } from './events'
// bind();

createApp(App).mount('#app')

async function setupApp() {
    const app = createApp(App)
    app.use(i18n)
    app.use(createPinia())

    const settingsStore = useSettingsStore()
    await settingsStore.loadPreferences()

    app.mount('#app')
}

setupApp()