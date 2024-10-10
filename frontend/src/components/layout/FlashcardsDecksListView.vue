<template>
    <div>
        <n-menu
            :options="menuOptions"
            :default-expanded-keys="[]"
        />
    </div>
</template>

<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import {
  DocumentHeaderFooter24Regular as FlashcardsIcon,
} from '@vicons/fluent'
import renderIcon from '@/utils/renderIcon'

const menuOptions: MenuOption[] = [
  {
    label: 'Deck 1',
    key: 'deck1',
    icon: renderIcon(FlashcardsIcon),
  },
  {
    label: 'Deck 2',
    key: 'deck2',
    icon: renderIcon(FlashcardsIcon),
    children: [
      {
        label: 'Subdeck Deck 3',
        key: 'deck3',
        icon: renderIcon(FlashcardsIcon),
      },
      {
        label: 'Subdeck Deck 4',
        key: 'deck4',
        icon: renderIcon(FlashcardsIcon),
      },
    ]
  },
]
</script>

<style lang="scss">

</style>