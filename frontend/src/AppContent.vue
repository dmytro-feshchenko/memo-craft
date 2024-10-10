<template>
  <n-spin :show="props.loading">
    <n-layout class="app-content-wrapper">
        <n-layout position="absolute" has-sider>
            <n-layout-sider
                content-style="padding: 24px;"
                :native-scrollbar="false"
                bordered
                class="app-menu"
            >
                <n-space vertical>
                    <n-input v-model:value="searchValue" size="small" type="text" placeholder="Search" />
                    <n-menu class="app-top-menu" :options="menuOptions" @update:value="handleUpdateValue" />
                    <div>
                        <div>Notes</div>
                        <div class="app-files-menu">
                            <file-tree-menu />
                        </div>
                    </div>
                    <div style="margin-top: 30px;">
                        <div>Flashcards</div>
                        <div class="app-files-menu">
                            <flashcards-decks-list-view />
                        </div>
                    </div>
                    <div style="margin-top: 30px;">
                        <div>Tags</div>
                        <div class="app-files-menu">
                            <tags-list-view />
                        </div>
                    </div>
                </n-space>
            </n-layout-sider>
            <n-layout content-style="padding: 24px;" :native-scrollbar="false">
                <navigation-tabs/>
                <div>Content here</div>
            </n-layout>
        </n-layout>
    </n-layout>
  </n-spin>
</template>

<script setup lang="ts">
import { NIcon, useMessage, NBadge } from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import { h, Component, ref } from 'vue'
import { RouterLink } from 'vue-router'
import {
    ArrowTrending24Regular as DashboardIcon,
    Folder24Regular as FolderIcon,
//   PersonOutline as PersonIcon,
//   WineOutline as WineIcon
} from '@vicons/fluent'
import FileTreeMenu from '@/components/menus/FileTreeMenu.vue'
import FlashcardsDecksListView from '@/components/layout/FlashcardsDecksListView.vue'
import TagsListView from '@/components/layout/TagsListView.vue'
import { InboxQueueIcon, NewCardsIcon, SettingsIcon } from '@/components/icons'
import NavigationTabs from '@/components/NavigationTabs.vue'

const searchValue = ref('');

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const props = defineProps({
  loading: Boolean,
})

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        'a',
        {
          to: {
            name: 'dashboard',
            params: {
              lang: 'en-US'
            }
          }
        },
        { default: () => 'Dashboard' }
      ),
    key: 'dashboard',
    icon: renderIcon(DashboardIcon),
  },
  {
    label: () =>
      h(
        'a',
        {
          to: {
            name: 'flashcards',
            params: {
              lang: 'en-US'
            }
          }
        },
        { default: () => 'Due today' }
      ),
    key: 'flashcards',
    icon: renderIcon(InboxQueueIcon),
    extra:  () => { return h(NBadge, { value: 22 }) }
  },
  {
    label: () =>
      h(
        'a',
        {
          to: {
            name: 'flashcards-new',
            params: {
              lang: 'en-US'
            }
          }
        },
        { default: () => 'New cards' }
      ),
    key: 'flashcards-new',
    icon: renderIcon(NewCardsIcon),
    extra:  () => { return h(NBadge, { value: 5 }) }
  },
  {
    label: () =>
      h(
        'a',
        {
          to: {
            name: 'settings',
            params: {
              lang: 'en-US'
            }
          }
        },
        { default: () => 'Settings' }
      ),
    key: 'settings',
    icon: renderIcon(SettingsIcon),
  },
]

const handleUpdateValue = (key: string, item: MenuOption) => {
    // message.info(`[onUpdate:value]: ${JSON.stringify(key)}`)
    // message.info(`[onUpdate:value]: ${JSON.stringify(item)}`)
}
</script>

<style lang="scss">
.app-content-wrapper {
    height: 100vh;
    text-align: left;
}
.app-top-menu {
    .n-menu-item {
        height: 28px;
    }
    .n-menu-item-content {
        padding-left: 0 !important;
        padding-right: 0;
        &:before {
            left: 0;
            right: 0;
        }
    }
}
.app-files-menu {
    margin-left: -32px;
}
</style>