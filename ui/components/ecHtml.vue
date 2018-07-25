<template>
    <div class="h" v-html="param.html"></div>
</template>

<script>
    export default {
        props: ['param']
    }
</script>

<style scoped>
    div.h {
        width: 100%;
        overflow: hidden;
    }
</style>