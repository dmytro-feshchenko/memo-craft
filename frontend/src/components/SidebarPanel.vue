<template>
    <div class="sidebar-panel">
        <div class="top">
            <ul class="options">
                <li><FolderIcon class="icon" /></li>
                <li><MagnifyingGlassIcon class="icon" /></li>
            </ul>
        </div>
        <div class="bottom">
            <ul class="options">
                <li><WrenchScrewdriverIcon class="icon" /></li>
            </ul>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { WrenchScrewdriverIcon } from '@heroicons/vue/24/solid'
import { MagnifyingGlassIcon } from '@heroicons/vue/24/solid'
import { FolderIcon } from '@heroicons/vue/24/solid'
</script>

<style scoped>
.sidebar-panel {
    text-align: center;
    display: flex;
    flex-direction: column;
    height: 100%;
    justify-content: space-between;
}
.options {
    list-style: none;
    list-style-type: none;
    padding: 0;
}
.options li {
    padding-top: 10px;
    cursor: pointer;
}
.icon {
    width: 24px;
}
</style>