<template>
    <card>
        <Form :label-width="80" style="width:400px;margin: auto">
            <FormItem prop="Username" label="用户名">
                <Input v-model="reg.Username" placeholder="请输入用户名">
                <span slot="prepend"><Icon :size="16" type="person"></Icon></span></Input>
                <div class="ivu-form-item-error-tip" v-if="em.Username!=''">{{em.Username}}</div>
            </FormItem>
            <FormItem prop="Email" label="邮箱">
                <Input v-model="reg.Email" placeholder="请输入邮箱">
                <span slot="prepend"><Icon :size="16" type="email"></Icon></span></Input>
                <div class="ivu-form-item-error-tip" v-if="em.Email!=''">{{em.Email}}</div>
            </FormItem>
            <FormItem prop="Mobile" label="手机号">
                <Input v-model="reg.Mobile" placeholder="手机号">
                <span slot="prepend"><Icon :size="16" type="iphone"></Icon></span></Input>
                <div class="ivu-form-item-error-tip" v-if="em.Mobile!=''">{{em.Mobile}}</div>
            </FormItem>
            <FormItem prop="Password" label="密码">
                <Input type="password" v-model="reg.Password" placeholder="请输入密码">
                <span slot="prepend"><Icon :size="14" type="locked"></Icon></span> </Input>
                <div class="ivu-form-item-error-tip" v-if="em.Password!=''">{{em.Password}}</div>
            </FormItem>
            <FormItem prop="CaptchaVal" label="认证码">
                <Input type="text" v-model="reg.CaptchaVal" placeholder="请输入认证码">
                <span slot="prepend"><Icon :size="14" type="flash"></Icon></span></Input>
                <img :src="authImg" @click="loadCaptcha"/>
                <div class="ivu-form-item-error-tip" v-if="em.CaptchaVal!=''">{{em.CaptchaVal}}</div>
            </FormItem>
            <FormItem>
                <Button type="primary" long size="large" @click="registerSubmit">注册</Button>
                <div style="float: right">
                    <router-link to="/p/user/login">登录</router-link>
                    <router-link to="/p/user/forget">忘记密码</router-link>
                </div>
            </FormItem>
        </Form>
    </card>
</template>

<script>import Cookies from 'js-cookie';
//iview validate bug,不能第二次验证
export default {
    data() {
        return {
            authImg: "",
            reg: {},

            em: {Username: '', Email: '', Mobile: '', Password: '', CaptchaVal: ''},
        }
    }, created() {
        this.loadCaptcha()
    }, methods: {
        validate() {
            var un = this.reg.Username;
            if (!un || un.length > 32 || un.length < 5) {
                this.em.Username = '用户名长度5~32';
                return false;
            } else {
                this.em.Username = '';
            }
            var e = this.reg.Email;
            if (!e || e.indexOf('@') < 0) {
                this.em.Email = '邮箱格式错误';
                return false;
            } else {
                this.em.Email = '';
            }
            e = this.reg.Mobile;
            if (!e || e.length != 11) {
                this.em.Mobile = '手机号格式错误';
                return false;
            } else {
                this.em.Mobile = '';
            }
            e = this.reg.Password;
            if (!e || e.length < 3) {
                this.em.Password = '密码长度大于3';
                return false;
            } else {
                this.em.Password = '';
            }
            e = this.reg.CaptchaVal;
            if (!e || e.length != 6) {
                this.em.CaptchaVal = '认证码长度6';
                return false;
            } else {
                this.em.CaptchaVal = '';
            }
            return true;
        },
        loadCaptcha() {
            this.ajax('/gk-user/CaptchaNew', {}, (r, th) => {
                th.CaptchaId = r.Result;
                th.authImg = "/api/gk-user/Captcha?t=" + th.CaptchaId;
            });
        },
        registerSubmit() {
            if (!this.validate()) {
                return
            }
            this.reg["CaptchaId"] = this.CaptchaId;
            this.ajax('/gk-user/Register', this.reg, function (r, th) {
                if (r.Status && (!r.Status.Code || r.Status.Code == 0)) {
                    Cookies.set('webToken', r.Result);
                    window.gk.user = r.Info;
                    window.gk.login = true;
                    th.$Modal.success({
                        title: "", content: "注册成功",
                        onOk() {
                            vm.$router.replace('/')
                        }
                    });
                } else {
                    th.regCap = "";
                    th.reg.CaptchaVal = "";
                    if (r.Status.Code == 8) {
                        th.em.CaptchaVal = r.Status.Msg;
                    }
                    if (r.Status.Code == 1000) {
                        th.em.Username = r.Status.Msg;
                    }
                    if (r.Status.Code == 1002) {
                        th.em.Email = r.Status.Msg;
                    }
                    th.CaptchaId = r.Result;
                    th.authImg = "/api/gk-user/Captcha?t=" + th.CaptchaId;
                }
            });

        },
    }
}
</script>
