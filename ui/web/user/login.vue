<template>
    <card>
        <div style="width:300px;margin: auto">
            <div class="login-con">
                <Card :bordered="false">
                    <p slot="title">
                        <Icon type="log-in"></Icon>
                        欢迎登录
                    </p>
                    <div class="form-con">
                        <Form :model="form" ref="form">
                            <FormItem prop="Username" :rules="rules.Username" label="用户名" :error="err.Username">
                                <Input v-model="form.Username" placeholder="请输入用户名">
                                <span slot="prepend">
                                    <Icon :size="16" type="person"></Icon>
                                </span>
                                </Input>
                            </FormItem>
                            <FormItem prop="Password" :rules="rules.Password" label="密码" :error="err.Password">
                                <Input type="password" v-model="form.Password" placeholder="请输入密码">
                                <span slot="prepend">
                                    <Icon :size="14" type="locked"></Icon>
                                </span>
                                </Input>
                            </FormItem>
                            <FormItem prop="password" v-if="authImg!=''" :error="err.Captcha">
                                <Input type="text" v-model="form.Digits" placeholder="请输入认证码">
                                </Input>
                                <img :src="authImg" @click="loadCaptcha"/>
                            </FormItem>
                            <FormItem>
                                <Button @click="handleSubmit" type="primary" long>登录</Button>
                            </FormItem>
                        </Form>
                    </div>
                </Card>
            </div>
        </div>
    </card>
</template>


<script>
    import Cookies from 'js-cookie';

    export default {
        data() {
            return {
                authImgUrl: '/api/gk-user/Captcha?t=',
                captchaNew: '/gk-user/CaptchaNew',
                authImg: "",
                form: {Captcha: "", Username: '', Digits: ''},
                err: {Username: "", Password: ""},
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
            Cookies.remove("webToken");
            var un = Cookies.get("user");
            if (un && un != "" && un.length > 1) {
                this.form.Username = un;
            }
        },
        methods: {
            handleSubmit() {
                for (var k in this.err) {
                    this.err[k] = "";
                }
                this.$refs.form.validate((valid) => {
                    if (valid) {
                        this.ajax('/gk-user/Login', this.form, (r, th) => {
                            if (r.code == 0) {
                                gk.user = r.result[0];
                                gk.login = true;
                                Cookies.set('webToken', gk.user.Token, {expires: 365});
                                if (window.goUrl == "" || window.goUrl == window.location.pathname) {
                                    window.goUrl = '/'
                                }
                                vm.$router.push({path: window.goUrl});
                                return;
                            }
                            if (r.result.length >= 2) {
                                th.form.Captcha = r.result[1];
                                th.authImg = th.authImgUrl + r.result[1];
                            } else {
                                th.authImg = ""
                            }
                            if (r.code == 1004) {
                                th.form.Password = "";
                            }
                            for (var k in th.err) {
                                th.err[k] = "";
                            }
                            th.err[r.result[0]] = r.msg;
                        });
                    }
                });
            },
            loadCaptcha() {
                let th = this;
                this.ajax(this.captchaNew, {}, function (r) {
                    th.CaptchaId = r.result[0];
                    th.authImg = th.authImgUrl + th.CaptchaId;
                });
            }
        }
    };
</script>