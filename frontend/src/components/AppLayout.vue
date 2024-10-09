<template>
    <div class="container">
        <div class="primary-menu">
            <SidebarPanel/>
        </div>
        <div class="secondary-menu">
            <div>Menu here 2</div>
        </div>
        <div class="content-wrapper">
            <CardPreview/>
            <InfoPanel/>
            <TabsPanel/>
            <FileEditor/>
        </div>
    </div>
</template>

<script lang="ts" setup>
import SidebarPanel from './SidebarPanel.vue'
import FileEditor from './FileEditor.vue'
import TabsPanel from './TabsPanel.vue'
import InfoPanel from './InfoPanel.vue'
import CardPreview from '../modules/flashcards/components/CardPreview.vue' 
</script>

<style scoped>
.container {
    display: flex;
    width: 100%;
    height: 100vh;
}
.primary-menu {
    width: 50px;
    background-color: var(--top-menu-bg);
    border-top: 1px solid var(--top-menu-border-color);
    border-right: 1px solid var(--top-menu-border-color);
    text-align: left;
    margin-top: 30px;
}
.secondary-menu {
    width: 200px;
    background-color: var(--secondary-menu-bg);
    border-right: 1px solid var(--secondary-menu-border-color);
    padding-top: 30px;
}
.content-wrapper {
    flex-grow: 1;
    text-align: left;
}
</style>