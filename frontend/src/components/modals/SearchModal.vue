<template>
    <n-modal v-model:show="showModal">
        <n-card
            title="Search notes and flashcards"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
            class="modal-wrap"
        >
            <n-space :size="10" :wrap="false" :wrap-item="false" align="center" vertical>
                <n-input v-model:value="searchValue" size="small" type="text" placeholder="Search" />
                <n-space :size="5" :wrap="false" :wrap-item="false" align="center">
                    <n-empty description="Notes not found"/>
                    <n-divider vertical />
                    <n-empty description="Flashcards not found"/>
                </n-space>
            </n-space>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// TODO: use searchStore.globalSearchEnabled
const searchValue = ref('');
const showModal = ref(false)
</script>

<style scoped>
.modal-wrap {
    width: 600px;
    position: fixed;
    top: 20px;
    left: 0;
    right: 0;
}
</style>