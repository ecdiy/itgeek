import Vue from 'vue';
import iView from 'iview';
import {router} from './libs/router/index';
import App from './libs/frame/app.vue';
import axios from "axios";
import Cookies from 'js-cookie';
import 'iview/dist/styles/iview.css';

Vue.use(iView);


Vue.prototype.ajax = function (url, p, fun) {
    let th = this;
    var pp = p ? p : {}
    pp["siteId"] = window.gk.siteId;
    axios.post('/api' + url, pp).then(function (r) {
        for (var k in r.data) {
            if (th.hasOwnProperty(k))
                th[k] = r.data[k]
        }
        if (th.hasOwnProperty("loading")) {
            th.loading = false;
        }
        if (r.data.hasOwnProperty("gkUser")) {
            gk.user = r.data.gkUser;
            vm.$emit("data", window.gk);
        }
        if (fun && typeof(fun) == 'function') {
            fun(r.data, th)
        }
    }).catch((err) => {
        if (err.response && err.response.status == 401) {
            window.goUrl = window.location.hash
            window.gk.login = false;
            window.gk.user = {};
            Cookies.remove('token');
            th.$router.replace('/p/login')
            vm.$emit("data", window.gk)
        } else {
            console.log(err)
        }
    })
}

window.goUrl = '/';
window.gk = {login: false, user: {}, siteId: 0};
Cookies.set("ua", "web");

Vue.component('go', function (resolve) {
    require(['../components/go.vue'], resolve)
})
window.vm = new Vue({
    el: '#app', router, render: h => h(App),
});
