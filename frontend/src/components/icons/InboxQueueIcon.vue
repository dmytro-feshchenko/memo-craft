<script lang="ts" setup>
const props = defineProps({
    size: {
        type: [String],
        default: "24px",
    },
    fill: {
        type: String,
        default: "none",
    },
})
</script>

<template>
    <svg :width="size" :height="size" viewBox="0 0 24 24" :fill="fill" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <rect width="24" height="24" fill="white"></rect> <path d="M3 15L3.92359 3.91695C3.96678 3.39866 4.40004 3 4.92013 3H19.0799C19.6 3 20.0332 3.39866 20.0764 3.91695L21 15M3 15V20C3 20.5523 3.44772 21 4 21H20C20.5523 21 21 20.5523 21 20V15M3 15H9C9 16 9.6 18 12 18C14.4 18 15 16 15 15H21" stroke="#000000" stroke-linejoin="round"></path> </g></svg>
</template>

<style lang="scss" scoped>
</style>