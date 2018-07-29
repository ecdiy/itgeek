<template>
    <div>
        <card>
            <Breadcrumb>
                <BreadcrumbItem>
                    <router-link to="/">首页</router-link>
                </BreadcrumbItem>
                <BreadcrumbItem> 日常任务</BreadcrumbItem>
            </Breadcrumb>
        </card>
        <card>
            <div v-if="LoginAward==1">
                <h3> 每日登录奖励 {{today}}</h3>
                <div style="padding: 15px;">
                    <Button @click="laDo">领取 X 铜币</Button>
                </div>
                <div>已连续登录 {{loginDay}} 天</div>
            </div>

            <div v-if="LoginAward==0">
                每日登录奖励已领取
                <div style="padding: 15px;width: 100%">
                    <go to="my/balance,1" title="账户余额">查看我的账户余额</go>
                </div>
                <div>已连续登录 {{loginDay}} 天</div>
            </div>
        </card>
    </div>

</template>

<script>
    export default {
        data() {
            return {loginDay: 0, today: "", gk: window.gk, LoginAward: 2}
        }, created() {
            this.ajax('/gk-user/LoginAwardStatus')
        }, methods: {
            laDo() {
                this.ajax('/gk-user/LoginAwardDo', {}, function (r, th) {
                    window.gk.user.Score = r.Score;
                    window.gk.user.LoginAward = 0;
                    th.LoginAward = 0;
                })
            }
        }
    }
</script>
