<template>

    <div>

        <div class="box" v-if="LoginAward==1">
            <span style="font-size: 18px;padding: 15px 8px;display: block"> 每日登录奖励 {{today}}</span>
            <x-button @click.native="laDo">领取 X {{dw}}</x-button>
            <div style=" padding: 15px 8px;display: block">已连续登录 {{loginDay}} 天</div>
        </div>

        <div class="box" v-if="LoginAward==0">
            每日登录奖励已领取, {{gk.user.Score}}
        </div>
    </div>

</template>

<script>
    export default {
        data() {
            return {loginDay: 0, today: "", gk: window.gk, LoginAward: 2, dw: gk.site.dw}
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
