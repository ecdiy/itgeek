<template>
    <card>
        <div style="text-align: center">
            <h2>
                {{gk.site.SiteName}}
            </h2>
            <span>{{gk.site.Sign}}</span>
            <hr>
            <div style="padding: 15px 0">
                <Button size="large" @click="goReg">现在注册</Button>
                <span> 已注册用户请 <router-link to="/p/user/login">登录</router-link></span>
            </div>
        </div>
    </card>
</template>
<script>
    export default {
        data() {
            return {gk: window.gk}
        },
        methods: {
            goReg: () => {
                vm.$router.replace('/p/user/register')
            }
        }
    }
</script>
<style scoped>
    span {
        clear: both;
        padding: 10px 0;
        display: block
    }
</style>