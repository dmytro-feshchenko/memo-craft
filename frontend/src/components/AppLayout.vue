<template>
    <div class="container">
        <div class="primary-menu">
            <SidebarPanel/>
        </div>
        <div class="secondary-menu">
            <div>Menu here 2</div>
        </div>
        <div class="content-wrapper">
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
</script>

<style scoped>
.container {
    display: flex;
    width: 100%;
    height: 100vh;
}
.primary-menu {
    width: 50px;
    background-color: rgba(49, 54, 60, 1);
    text-align: left;
    padding-top: 30px;
}
.secondary-menu {
    width: 200px;
    background-color: rgba(49, 54, 60, 1);
    border-right: 1px solid rgba(63, 69, 75, 1);
    padding-top: 30px;
}
.content-wrapper {
    flex-grow: 1;
    text-align: left;
}
</style>