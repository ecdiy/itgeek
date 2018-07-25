<template>
    <card>
        <topic-list :list="FavList"></topic-list>
    </card>
</template>

<script>
    import topicList from '../topic/topicList'

    export default {
        components: {topicList},
        data() {
            return {FavList: [], page: 1}
        },
        created() {
            this.ajax('/gk-topic/followList', {page: this.page})
        }
    }
</script>
