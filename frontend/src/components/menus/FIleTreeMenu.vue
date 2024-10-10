<template>
    <div>
        <n-menu
            :options="menuOptions"
            :default-expanded-keys="[]"
            accordion
        />
    </div>
</template>

<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import {
    Folder24Regular as FolderIcon,
} from '@vicons/fluent'
import renderIcon from '@/utils/renderIcon'

const menuOptions: MenuOption[] = [
  {
    label: 'Folder 1',
    key: 'folder1',
    icon: renderIcon(FolderIcon),
    children: [
      {
        label: 'file1.md',
        key: 'file1.md',
      },
      {
        label: 'file2.md',
        key: 'file2.md',
      },
    ]
  },
  {
    label: 'Folder 2',
    key: 'folder2',
    icon: renderIcon(FolderIcon),
    children: [
        {
            label: 'Folder 2',
            key: 'folder2',
            icon: renderIcon(FolderIcon),
            children: [
                {
                    label: 'file3.md',
                    key: 'file3.md',
                },
                {
                    label: 'file4.md',
                    key: 'file4.md',
                },
            ]
        },
    ]
  },
]
</script>

<style lang="scss">

</style>