<template>

    <card>
        <Breadcrumb>
            <BreadcrumbItem>
                <router-link to="/">首页</router-link>
            </BreadcrumbItem>
            <BreadcrumbItem> 日常任务</BreadcrumbItem>
        </Breadcrumb>

        <div v-if="LoginAward==1">
            <h1> 每日登录奖励 {{today}}</h1>
            <Button @click="laDo">领取 X 铜币</Button>
            <div>已连续登录 {{days}} 天</div>
        </div>

        <div v-if="LoginAward==0">
            每日登录奖励已领取, {{gk.user.Score}}
        </div>
    </card>

</template>

<script>
    export default {
        data() {
            return {days: 0, today: "", gk: window.gk, LoginAward: 2}
        }, created() {
            this.ajax('/gk-user/LoginAwardStatus')
        }, methods: {
            laDo() {
                this.ajax('/gk-user/LoginAwardDo', {}, function (r,th) {
                    window.gk.user.Score = r.Score;
                    window.gk.user.LoginAward = 0;
                    th.LoginAward = 0;
                })
            }
        }
    }
</script>
