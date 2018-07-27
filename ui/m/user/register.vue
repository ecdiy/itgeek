<template>
    <div>
        <x-input v-model="reg.Username" title="用户名" placeholder="请输入用户名" novalidate
                 :icon-type="icon.Username"></x-input>
        <x-input v-model="reg.Email" title="邮 箱" placeholder="请输入邮箱" novalidate :icon-type="icon.Email"></x-input>
        <x-input v-model="reg.Mobile" title="手机号" placeholder="请输入手机号" novalidate :icon-type="icon.Mobile"></x-input>
        <x-input v-model="reg.Password" title="密 码" type="password" placeholder="请输入密码"
                 novalidate :icon-type="icon.Password"></x-input>
        <x-input v-model="reg.CaptchaVal" title="认证码" placeholder="请输入认证码" novalidate
                 :icon-type="icon.CaptchaVal"></x-input>
        <img :src="authImg" @click="loadCaptcha"/>
        <x-button @click.native="registerSubmit">注册</x-button>
    </div>
</template>

<script>import {AlertPlugin, ToastPlugin} from 'vux'
import Cookies from 'js-cookie';
import Vue from 'vue';

Vue.use(ToastPlugin)
Vue.use(AlertPlugin)

export default {
    data() {
        return {authImg: "", reg: {}, regUn: "", regEm: "", regCap: "", icon: {Username: ''}}
    }, created() {
        this.loadCaptcha()
    }, methods: {
        validate() {
            var un = this.reg.Username;
            if (!un || un.length > 32 || un.length < 5) {
                this.icon.Username = 'error';
                return false;
            } else {
                this.icon.Username = 'success';
            }
            var e = this.reg.Email;
            if (!e || e.indexOf('@') < 0) {
                this.icon.Email = 'error';
                return false;
            } else {
                this.icon.Email = 'success';
            }
            e = this.reg.Mobile;
            if (!e || e.length != 11) {
                this.icon.Mobile = 'error';
                return false;
            } else {
                this.icon.Mobile = 'success';
            }
            e = this.reg.Password;
            if (!e || e.length < 3) {
                this.icon.Password = 'error';
                return false;
            } else {
                this.icon.Password = 'success';
            }
            e = this.reg.CaptchaVal;
            if (!e || e.length != 6) {
                this.icon.CaptchaVal = 'error';
                return false;
            } else {
                this.icon.CaptchaVal = 'success';
            }
            return true;
        },
        loadCaptcha() {
            this.ajax('/gk-user/CaptchaNew', {}, function (r, th) {
                th.CaptchaId = r.Result;
                th.authImg = "/api/gk-user/Captcha?t=" + th.CaptchaId;
            });
        },
        registerSubmit() {
            if (!this.validate()) {
                return
            }
            this.reg["CaptchaId"] = this.CaptchaId;
            this.ajax('/gk-user/Register', this.reg, (r, th) => {
                if (r.Status && !r.Status.Code) {
                    Cookies.set('h5Token', r.Result, {expires: 365});
                    window.gk.user = r.Info;
                    window.gk.login = true;
                    vm.$emit("data", window.gk);

                    th.$vux.alert.show({
                        title: '注册成功',
                        onHide() {
                            window.location = "#/";
                        }
                    });
                } else {
                    th.reg.CaptchaVal = "";
                    th.CaptchaId = r.Result;
                    th.authImg = "/api/gk-user/Captcha?t=" + th.CaptchaId;
                    th.$vux.toast.text(r.Status.Msg, 'bottom')
                    if (r.Status.Code == 8) {
                        th.icon.CaptchaVal = 'error';
                    }
                    if (r.Status.Code == 1000) {
                        th.icon.Username = 'error';
                    }
                    if (r.Status.Code == 1002) {
                        th.icon.Email = 'error';
                    }
                }
            });
        },
    }
}
</script>
