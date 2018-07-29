<template>
    <card>
        <div style="width:300px;margin: auto">
            <div class="login-con">
                <Card :bordered="false">
                    <span slot="title">
                        <Icon type="log-in"></Icon>
                        {{title}}
                    </span>
                    <div class="form-con">
                        <Form :model="form" ref="form">
                            <FormItem prop="SiteAdminUser" :rules="rules.SiteAdminUser" label="用户名"
                                      :error="err.SiteAdminUser">
                                <Input v-model="form.SiteAdminUser" placeholder="请输入用户名">
                                <span slot="prepend">
                                    <Icon :size="16" type="person"></Icon>
                                </span>
                                </Input>
                            </FormItem>
                            <FormItem prop="SiteAdminPass" :rules="rules.SiteAdminPass" label="密码"
                                      :error="err.SiteAdminPass">
                                <Input type="password" v-model="form.SiteAdminPass" placeholder="请输入密码">
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
                title: '', url: '',
                authImgUrl: '/api/gk-user/Captcha?t=',
                captchaNew: '/gk-user/CaptchaNew',
                authImg: "",
                form: {Captcha: "", SiteAdminUser: '', Digits: ''},
                err: {SiteAdminUser: "", SiteAdminPass: ""},
                rules: {
                    SiteAdminUser: [
                        {required: true, min: 3, message: '账号不能为空，长度至少5位', trigger: 'blur'}
                    ],
                    SiteAdminPass: [
                        {required: true, min: 6, message: '密码不能为空，长度至少6位', trigger: 'blur'}
                    ]
                }
            };
        },
        created() {
            Cookies.remove("webGeekAdmin");
            this.ajax('/gk-admin/userStatus', {}, (r, th) => {
                if (!r.login && !r.user && !r.pass) {
                    th.title = "设置管理员账号"
                    th.url = 'userInit'
                } else {
                    th.title = '欢迎登录'
                    th.url = 'login'
                }
            });

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
                        this.ajax('/gk-admin/' + this.url, this.form, (r, th) => {
                            if (r.code == 0) {
                                Cookies.set('webGeekAdmin', r.result[0], {expires: 365});
                                gk.login = true;
                                vm.$router.push({path: '/'});
                            } else {
                                gk.login = false;
                                th.$Notice.error({title: r.msg});
                            }
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