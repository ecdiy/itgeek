import Vue from 'vue';
import iView from 'iview';
import {router} from './libs/router/index';
import App from './libs/frame/app.vue';
import axios from "axios";
import Cookies from 'js-cookie';
import 'iview/dist/styles/iview.css';

Vue.use(iView);

// String.prototype.replaceAll = function (s1, s2) {
//     return this.replace(new RegExp(s1, "gm"), s2);
// }

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
        if (th.hasOwnProperty("loading")) {
            th.loading = false;
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
            Cookies.remove('token');
            th.$router.replace('/p/user/login')
            vm.$emit("data", window.gk)
        } else {
            console.log(err)
        }
    })
}
Vue.component('go', function (resolve) {
    require(['../components/go.vue'], resolve)
})
Vue.component('gk-body', function (resolve) {
    require(['./libs/frame/gkBody.vue'], resolve)
})
Vue.component('base-page', function (resolve) {
    require(['../components/base-table.vue'], resolve)
})

Vue.component('user-info', function (resolve) {
    require(['./user/sm-user-info.vue'], resolve)
})

Vue.component('login', function (resolve) {
    require(['./user/sm-login.vue'], resolve)
})


Vue.component('gk-img', function (resolve) {
    require(['../components/ecImg.vue'], resolve)
})

Vue.component('gk-html', function (resolve) {
    require(['../components/ecHtml.vue'], resolve)
})

window.goUrl = '/';

window.gk = {login: false, user: {}, site: {}, siteId: 0};

window.vm = new Vue({el: '#app', router, render: h => h(App)});

window.vgo = (url, o) => {
    if (o)
        vm.$emit("msgClick", o.parentNode.getAttribute("msg-id"))
    vm.$router.replace(url)
};

Cookies.set("ua", "web");