<style scoped>
    .none {
        display: none;
    }
</style>
<template>
    <div @keydown.enter="handleSubmit" class="cell">

        <!--<div class="form-con">-->
        <!--<Form :model="form" ref="form">-->
        <!--<FormItem prop="Username" :rules="rules.Username" label="用户名" :error="err.Username">-->
        <!--<Input v-model="form.Username" placeholder="请输入用户名">-->
        <!--<span slot="prepend">-->
        <!--<Icon :size="16" type="person"></Icon>-->
        <!--</span>-->
        <!--</Input>-->

        <x-input title="用户名" :show-clear="true" type="text" v-model="form.Username" :placeholder="'请输入用户名'"></x-input>

        <!--<FormItem prop="Password" :rules="rules.Password" label="密码" :error="err.Password">-->
        <!--<Input type="password" v-model="form.Password" placeholder="请输入密码">-->
        <!--<span slot="prepend">-->
        <!--<Icon :size="14" type="locked"></Icon>-->
        <!--</span>-->
        <!--</Input>-->
        <x-input title="密 码" :placeholder="'请输入密码'"
                 type="password" :maxlength="32" v-model="form.Password"></x-input>
        <!--</FormItem>-->
        <!--<FormItem prop="password" :class="captchaCls" :error="err.Captcha">-->
        <!--<Input type="text" v-model="form.Digits" placeholder="请输入认证码">-->
        <!--</Input>-->
        <!--<img :src="authImg" @click="loadCaptcha"/>-->
        <!--</FormItem>-->
        <!--<FormItem>-->
        <!--<cube-button @click="handleSubmit">登录</cube-button>-->
        <x-button @click.native="handleSubmit" type="primary">登录</x-button>
        <!--<Button @click="handleSubmit" type="primary" long>登录</Button>-->
        <!--</FormItem>-->
        <!--</Form>-->
    </div>
</template>

<script>
    import Cookies from 'js-cookie';

    export default {
        // props: ['cookieTokenName', 'authImgUrl', 'captchaNew', 'actionUrl'],
        data() {
            return {

                captchaCls: "none", authImg: "", form: {Captcha: ""},
                rules: {
                    Username: [
                        {required: true, min: 3, message: '账号不能为空，长度至少5位', trigger: 'blur'}
                    ],
                    Password: [
                        {required: true, min: 6, message: '密码不能为空，长度至少6位', trigger: 'blur'}
                    ]
                }
            };
        },
        created() {
            Cookies.remove('token');
            var un = Cookies.get("user");
            if (un && un != "" && un.length > 1) {
                this.form.Username = un;
            }
        },
        methods: {
            loadCaptcha() {
                this.ajax(this.captchaNew, {}, function (r, th) {
                    th.CaptchaId = r.Result;
                    th.authImg = th.authImgUrl + th.CaptchaId;
                });
            },
            handleSubmit() {
                //  this.$createToast({type: 'error', txt: 'Timeout'}).show()


                this.ajax('/gk-user/Login', this.form, function (r, th) {
                    if ((!r.Status.Code || r.Status.Code == 0) && r.Result.length > 0) {
                        Cookies.set('token', r.Result);
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
                            th.form.Captcha = r.Result;
                            th.form.Digits = '';
                            th.authImg = th.authImgUrl + th.form.Captcha;
                        }
                    }
                });

            }
        }
    };
</script>