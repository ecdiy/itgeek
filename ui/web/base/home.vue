<template>
    <gk-body ec-key="Index">
        <bm-list slot="left"></bm-list>
        <div slot="right">
            <count-info></count-info>
        </div>
    </gk-body>
</template>

<script>
    import bmList from '../topic/bm-list';

    import countInfo from '../topic/smCountInfo';

    export default {
        components: {
            bmList, countInfo
        }
    }
</script>
