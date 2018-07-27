<style scoped>
    .none {
        display: none;
    }
</style>
<template>
    <div @keydown.enter="handleSubmit" style="width:300px;margin: auto">
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
                        <FormItem prop="password" :class="captchaCls" :error="err.Captcha">
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
</template>

<script>
    import Cookies from 'js-cookie';

    export default {
        props: ['cookieTokenName', 'authImgUrl', 'captchaNew', 'actionUrl'],
        data() {
            return {
                captchaCls: "none", authImg: "", form: {Captcha: ""}, err: {Username: "", Password: ""},
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
            Cookies.remove(this.cookieTokenName);
            var un = Cookies.get("user");
            if (un && un != "" && un.length > 1) {
                this.form.Username = un;
            }
        },
        methods: {
            loadCaptcha() {
                let th = this;
                this.ajax(this.captchaNew, {}, function (r) {
                    th.CaptchaId = r.Result;
                    th.authImg = th.authImgUrl + th.CaptchaId;
                });
            },
            handleSubmit() {
                for (var k in this.err) {
                    this.err[k] = "";
                }
                this.$refs.form.validate((valid) => {
                    if (valid) {
                        this.ajax(this.actionUrl, this.form, (r, th) => {
                            if ((!r.Status.Code || r.Status.Code == 0) && r.Result.length > 0) {
                                Cookies.set(th.cookieTokenName, r.Result, {expires: 365});
                                th.$emit('on-login', r.Info);
                            } else {
                                if (r.Status.Code == 8) {
                                    th.captchaCls = '';
                                } else {
                                    th.form.Password = "";
                                    for (var k in th.err) {
                                        th.err[k] = "";
                                    }
                                }
                                if (r.Result != '') {
                                    th.err[r.Result] = r.Status.Msg;
                                }
                                if (r.Result == 'Captcha') {
                                    th.captchaCls = '';
                                    th.form.Captcha = "";
                                    th.form.Digits = '';
                                    th.form.Captcha = r.Param.Val;
                                    th.authImg = th.authImgUrl + th.form.Captcha;
                                }
                            }
                        });
                    }
                });
            }
        }
    }
</script>
