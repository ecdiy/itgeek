<template>
    <div @keydown.enter="handleSubmit" class="box">
        <x-input title="用户名" :show-clear="true" type="text" v-model="form.Username" :placeholder="'请输入用户名'"></x-input>
        <x-input title="密 码" :placeholder="'请输入密码'"
                 type="password" :maxlength="32" v-model="form.Password"></x-input>

        <x-input type="text" v-model="form.Digits" placeholder="请输入认证码">
            <img :src="authImg" slot="right-full-height" @click="loadCaptcha"/>
        </x-input>

        <x-button @click.native="handleSubmit" type="primary">登录</x-button>
    </div>
</template>

<script>
    import Cookies from 'js-cookie';

    export default {
        data() {
            return {
                captchaCls: "none", authImg: "", form: {Captcha: ""},
            };
        },
        created() {
            Cookies.remove('h5Token');
            var un = Cookies.get("user");
            if (un && un != "" && un.length > 1) {
                this.form.Username = un;
            }
        },
        methods: {
            loadCaptcha() {
                this.ajax(this.captchaNew, {}, function (r, th) {
                    th.CaptchaId = r.result;
                    th.authImg = th.authImgUrl + th.CaptchaId;
                });
            },
            handleSubmit() {
                this.ajax('/gk-user/Login', this.form, function (r, th) {
                    if ((!r.Code || r.Status.Code == 0) && r.result.length > 0) {
                        Cookies.set('h5Token', r.result, {expires: 365});
                        window.gk.user = r.Param;
                        window.gk.login = true;
                        vm.$emit("data", window.gk);
                        window.location = window.goUrl
                    } else {
                        th.form.Password = "";
                        th.form.Captcha = "";
                        for (var k in th.err) {
                            th.err[k] = "";
                        }
                        th.err[r.Status.Target] = r.Status.Msg;
                        if ('Captcha' == r.Status.Target) {
                            th.captchaCls = '';
                            th.form.Captcha = r.result;
                            th.form.Digits = '';
                            th.authImg = th.authImgUrl + th.form.Captcha;
                        }
                    }
                });
            }
        }
    };
</script>