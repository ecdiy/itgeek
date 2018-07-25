<template>
    <card>
        <Form :label-width="80">
            <FormItem label="注册会员">
                <span>{{info.UserCount}}</span>
            </FormItem>
            <FormItem label="主题">
                <span>{{info.TopicCount}}</span>
            </FormItem>
            <FormItem label="回复">
                <span>{{info.ReplyCount}}</span>
            </FormItem>
        </Form>
    </card>
</template>

<script>
    export default {
        data() {
            return {info: {}}
        },
        mounted() {

            this.ajax("/gk-user/CountInfo", {}, function (r,th) {
                th.info = r;
            })
        }
    }
</script>

<style scoped>
    .ivu-form-item {
        margin-bottom: 0
    }
</style>