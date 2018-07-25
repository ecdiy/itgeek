<template>
    <div>
        <card>
            <p slot="title">
                <Icon type="stats-bars"></Icon>
                今日阅读排行
            </p>
            <p v-for="(it,i ) in today">
                <router-link :to="'/p/topic/detail,'+it.Id+','+it.UserId">{{i+1}}.{{it.Title}}</router-link>
            </p>
        </card>

    </div>
</template>

<script>
    export default {
        data() {
            return {today: []}
        },
        created() {
            this.ajax("/gk-topic/topic/hot");
        }
    }
</script>
