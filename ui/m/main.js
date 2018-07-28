import Vue from 'vue';
import {router} from './libs/router/index';
import App from './base/app.vue';
import axios from "axios";
import Cookies from 'js-cookie';

import {XInput, XButton, Group, Cell, Flexbox, FlexboxItem, ConfirmPlugin} from 'vux'

Vue.component('x-input', XInput)
Vue.component('x-button', XButton)
Vue.component('group', Group)
Vue.component('cell', Cell)
Vue.component('flexbox', Flexbox)
Vue.component('flexbox-item', FlexboxItem)
Vue.component('go', function (resolve) {
    require(['../components/go.vue'], resolve)
})
Vue.component('card', function (resolve) {
    require(['../components/mCard.vue'], resolve)
})
Vue.use(ConfirmPlugin)

window.avatarVer = (new Date()).getTime()
Vue.prototype.avatar = function (id) {
    return '/avatar/' + gk.siteId + '/' + id + '/48.png?t=' + window.avatarVer
}

Vue.prototype.ajax = function (url, p, fun) {
    let th = this;
    axios.post('/api' + url, p ? p : {}).then(function (r) {
        for (var k in r.data) {
            if (th.hasOwnProperty(k))
                th[k] = r.data[k]
        }
        if (fun && typeof(fun) == 'function') {
            fun(r.data, th)
            vm.$emit("data", window.gk)
        }
    }).catch((err) => {
        if (err.response && err.response.status == 401) {
            window.goUrl = window.location.hash
            window.gk.login = false;
            window.gk.user = {};
            Cookies.remove('h5Token');
            th.$router.replace('/p/user/login')
            vm.$emit("data", window.gk)
        } else {
            console.log(err)
        }
    })
}

window.goUrl = '/';
window.gk = {login: false, user: {}, site: {}, siteId: 0};

window.vm = new Vue({el: '#app', router: router, render: h => h(App)});

