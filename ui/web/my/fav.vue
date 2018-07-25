<template>
    <Card>
        <topic-list :list="FavList"></topic-list>
    </Card>

</template>

<script>
    import topicList from '../topic/topicList'

    export default {
        components: {topicList},
        data() {
            return {FavList: [], page: 1}
        },
        created() {
            this.ajax('/gk-topic/favList', {page: this.page})
        }
    }
</script>
