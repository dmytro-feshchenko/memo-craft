<template>
    <div class="tags-list-view">
        <n-menu
            :options="menuOptions"
            :default-expanded-keys="[]"
        />
    </div>
</template>

<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import { HashtagIcon } from '@/components/icons'
import renderIcon from '@/utils/renderIcon'

const menuOptions: MenuOption[] = [
  {
    label: 'Test tag 1',
    key: 'tag1',
    icon: renderIcon(HashtagIcon),
  },
  {
    label: 'Test tag 2',
    key: 'tag1',
    icon: renderIcon(HashtagIcon),
  },
  {
    label: 'Test tag 3',
    key: 'tag1',
    icon: renderIcon(HashtagIcon),
  },
]
</script>

<style lang="scss">
.tags-list-view {
    .n-menu .n-menu-item {
        height: 22px;
    }
    .n-menu-item-content__icon,
    .n-menu-item-content__icon svg {
        width: 12px;
        height: 12px;
        fill: #808080;
    }
}
</style>