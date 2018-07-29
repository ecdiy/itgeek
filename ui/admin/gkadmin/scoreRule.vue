<template>
    <div>
        <Form :labelWidth="140" style="width: 400px">
            <FormItem label="初始资本">
                <Input v-model="fm.Register"/>

            </FormItem>
            <FormItem label="创建主题">
                <Input v-model="fm.Topic"/>

            </FormItem>
            <FormItem label="回复">
                <Input v-model="fm.Reply"/>

            </FormItem>
            <FormItem label="收到回复">
                <Input v-model="fm.GetReply"/>

            </FormItem>
            <FormItem label="每日登录奖励最小值">
                <Input v-model="fm.LoginMin"/>
            </FormItem>
            <FormItem label="每日登录奖励最大值">
                <Input v-model="fm.LoginMax"/>
            </FormItem>
            <FormItem label="谢意">
                <Input v-model="fm.Thank"/>
            </FormItem>
            <FormItem>
                <Button @click="save">保存</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>
    export default {
        data() {
            return {fm: {}}
        }, methods: {
            save() {
                this.ajax('/gk-admin/baseSave', {info: JSON.stringify(this.fm), key: "ScoreRule"}, (o, th) => {
                    th.$Notice.success({
                        title: '设置保存成功'
                    });
                })
            }
        }, created() {
            this.ajax("/gk-admin/baseGet", {key: "ScoreRule"}, (o, t) => {
                if (o.baseInfo && o.baseInfo.length > 2) {
                    t.fm = JSON.parse(o.baseInfo)
                } else {
                    t.fm = {};
                }
                if (!t.fm.Thank) {
                    t.fm.Thank = -10
                }
                if (!t.fm.LoginMax) {
                    t.fm.LoginMax = 50
                }
                if (!t.fm.LoginMin) {
                    t.fm.LoginMin = 5
                }
                if (!t.fm.Reply) {
                    t.fm.Reply = -5
                }
                if (!t.fm.GetReply) {
                    t.fm.GetReply = 5
                }
                if (!t.fm.Topic) {
                    t.fm.Topic = -20
                }
                if (!t.fm.Register) {
                    t.fm.Register = 2000
                }
            })
        }
    }
</script>
