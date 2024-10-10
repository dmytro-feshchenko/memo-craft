<script lang="ts" setup>
import { h, onMounted, ref, watch } from 'vue'
// import AppLayout from './components/AppLayout.vue'
import { Environment, WindowSetDarkTheme, WindowSetLightTheme } from '@/../wailsjs/runtime/runtime'
import AppContent from './AppContent.vue'
import { useI18n } from 'vue-i18n'
import { useSettingsStore } from '@/stores/settings'
import { darkTheme, NButton, NSpace } from 'naive-ui'
import SearchModal from '@/components/modals/SearchModal.vue'

const settingsStore = useSettingsStore()
const i18n = useI18n()
const initializing = ref(true)
onMounted(async () => {
    try {
        initializing.value = true
        // load required stores

        // const env = await Environment()
    } finally {
        initializing.value = false
    }
})

// watch theme and dynamically switch
// watch(
//     () => settingsStore.isDark,
//     (isDark) => (isDark ? WindowSetDarkTheme() : WindowSetLightTheme()),
// )

// watch language and dynamically switch
watch(
    () => settingsStore.general.language,
    (lang) => (i18n.locale.value = settingsStore.currentLanguageCode),
)
</script>

<template>
  <n-config-provider
      :inline-theme-disabled="true"
      :locale="settingsStore.themeLocale"
      class="fill-height">
      <n-dialog-provider>
          <app-content :loading="initializing" />

          <!-- top modal dialogs here -->
          <search-modal />
      </n-dialog-provider>
  </n-config-provider>
</template>

<style scoped>
</style>
