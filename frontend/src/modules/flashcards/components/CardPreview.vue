<template>
    <div class="preview">
        <div class="action-bar">
                <button type="button" class="action"><PencilIcon class="action-icon" /></button>
                <button type="button" class="action"><TrashIcon class="action-icon" /></button>
                <button type="button" class="action"><EllipsisVerticalIcon class="action-icon" /></button>
            </div>
        <div class="preview-body">
            <div>Rendered markdown content here...</div>
            <hr>
            <div>Rendered next markdown content here...</div>
        </div>
        <div class="preview-footer">
            <div><PlusCircleIcon class="status-icon" />New card</div>
            <!-- <div>Source indicator</div> -->
        </div>
    </div>
</template>

<script setup lang="ts">
import { PencilIcon } from '@heroicons/vue/24/solid'
import { TrashIcon } from '@heroicons/vue/24/solid'
import { EllipsisVerticalIcon } from '@heroicons/vue/24/solid'
import { PlusCircleIcon } from '@heroicons/vue/24/solid'

</script>

<style scoped>
.preview {
    margin: 10px;
    background: var(--flashcard-overlay-color);
    /* background: rgba(240, 240, 240, 1); */
    border: 1px solid var(--flashcard-border-color);
    border-radius: 5px;
    position: relative;
    color: rgba(137, 137, 137, 1);

}
.preview-body {
    border-bottom: 1px solid var(--flashcard-border-color);
    border-radius: 5px;
    background: white;
    color: var(--flashcard-text-color);
}
.preview-body hr {
    border: 1px dashed var(--flashcard-border-color);
}
.preview-body div {
    padding: 12px 20px;
}
/* Action bar */
.action-bar {
    border: 1px solid var(--flashcard-border-color);
    position: absolute;
    top: -5px;
    right: 10px;
    border-radius: 5px;
    background-color: white;
}
.action {
    /* border-style: none; */
    background: none;
	border: none;
	cursor: pointer;
}
.action .action-icon {
    height: 16px;
    fill: var(--flashcard-action-text-color);
}
.action:hover {
    background-color: var(--flashcard-action-bg-color-hover);
}
.action:hover .action-icon {
    fill: var(--flashcard-action-text-color-hover);
}
/* Footer */
.preview-footer {
    padding: 5px 15px;
    font-size: 12px;
    line-height: 24px;
}
.status-icon {
    height: 14px;
    fill: var(--flashcard-action-text-color);
}
</style>