<template>
    <div>
        <div class="top">
            <ul class="options">
                <li>1</li>
                <li>2</li>
                <li>3</li>
            </ul>
        </div>
        <div class="bottom">
            <ul class="options">
                <li><BeakerIcon class="size-6 text-blue-500" /></li>
            </ul>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { BeakerIcon } from '@heroicons/vue/24/solid'
</script>

<style scoped>
.options {
    list-style: none;
    list-style-type: none;
}
</style>