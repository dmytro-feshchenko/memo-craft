<template>
    <div class="container">
        <div class="primary-menu">
            <SidebarPanel/>
        </div>
        <div class="secondary-menu">
            <div>Menu here</div>
        </div>
        <div class="content-wrapper">
            <div>Content here</div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import SidebarPanel from './SidebarPanel.vue'
</script>

<style scoped>
.container {
    display: flex;
    width: 100%;
}
.primary-menu {
    width: 50px;
}
.secondary-menu {
    width: 200px;
}
.content-wrapper {
    flex-grow: 1;
}
</style>