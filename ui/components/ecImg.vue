<template>
    <div v-if="this.param && this.param.list">
        <div v-for="it in this.param.list">
            <a target="_blank" :href="it.href" v-if="it.href && it.href.length>1"> <img :src="it.url"/></a>
            <img v-else :src="it.url"/>
        </div>
    </div>
</template>

<script>
    export default {
        props: ['param']
    }
</script>

<style scoped>
    img {
        border: 0;
        width: 100%;
        overflow: hidden;
    }

    a {
        padding: 0;
        margin: 0
    }
</style>