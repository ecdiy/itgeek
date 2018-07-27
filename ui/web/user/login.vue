<template>
    <card>
        <login cookie-token-name="webToken" captcha-new="/gk-user/CaptchaNew"
               @on-login="handleLogin" action-url="/gk-user/Login" auth-img-url="/gk-user/Captcha?t="></login>
    </card>
</template>

<script>
    import login from '../../components/login';

    export default {
        components: {login},
        methods: {
            handleLogin(d) {
                window.gk.user = d;
                window.gk.login = true;
                if (window.goUrl == "" || window.goUrl == window.location.pathname) {
                    window.goUrl = '/'
                }
                vm.$router.replace(window.goUrl)
            }
        }
    };
</script>