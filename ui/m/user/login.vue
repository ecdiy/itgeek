<template>
    <div @keydown.enter="handleSubmit" class="box">
        <x-input :title='`<span style="${styleUsername}">用户名</span>`' :show-clear="true" type="text"
                 v-model="form.Username" :placeholder="'请输入用户名'"></x-input>
        <x-input :title='`<span style="${stylePassword}">密 码</span>`' :placeholder="'请输入密码'"
                 type="password" :maxlength="32" v-model="form.Password"></x-input>

        <x-input :title='`<span style="${styleCaptcha}">认证码</span>`' v-if="authImg!=''" type="text"
                 v-model="form.Digits"
                 placeholder="请输入认证码">
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
                authImg: "", form: {Captcha: "", Username: '', Digits: ''},
                authImgUrl: '/api/gk-user/Captcha?t=', err: ''
            };
        },
        created() {
            window.gk.login = false;
            Cookies.remove('h5Token');
            var un = Cookies.get("user");
            if (un && un != "" && un.length > 1) {
                this.form.Username = un;
            }
        },
        computed: {
            styleCaptcha() {
                return this.style('Captcha')
            }, styleUsername() {
                return this.style('Username')
            }, stylePassword() {
                return this.style('Password')
            }
        },
        methods: {
            style(n) {
                return n == this.err ? "color:red;" : '';
            },
            loadCaptcha() {
                this.ajax('/gk-user/CaptchaNew', {}, function (r, th) {
                    th.CaptchaId = r.result;
                    th.authImg = th.authImgUrl + th.CaptchaId;
                });
            },
            handleSubmit() {
                if (this.form.Username && this.form.Password && this.form.Username.length > 3 && this.form.Password.length > 3) {
                    this.ajax('/gk-user/Login', this.form, function (r, th) {
                        if (r.code == 0) {
                            gk.user = r.result[0];
                            gk.login = true;
                            Cookies.set('h5Token', gk.user.Token, {expires: 365});
                            vm.$router.push({path: '/'});
                            return;
                        }
                        th.err = r.result[0];
                        th.form.Digits = '';
                        if (r.result.length >= 2) {
                            th.form.Captcha = r.result[1];
                            th.authImg = th.authImgUrl + r.result[1];
                        } else {
                            th.authImg = ""
                        }
                        if (r.code == 1004) {
                            th.form.Password = "";
                        }
                    });
                }
            }
        }
    };
</script>