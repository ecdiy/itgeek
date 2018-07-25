import Vue from 'vue';
import VueRouter from 'vue-router';
import {routers} from './router';

Vue.use(VueRouter);


export const router = new VueRouter({mode: 'history', routes: routers});

router.beforeEach((to, from, next) => {
    // var fd = false;
    // var pos = 0;
    // for (var i = 0; i < routers.length; i++) {
    //     var c = routers[i].children;
    //     for (var j = 0; j < c.length; j++) {
    //         if (c[j].name == to.name) {
    //             window.document.title = '极客社区 - ' + c[j].title;
    //             fd = true;
    //             pos = i;
    //             break;
    //         }
    //     }
    // }
    // if (fd) {
    //     if (pos === 1 && !window.gk.login) {
    //         window.goUrl = '#' + to.path;
    //         next({path: '/user/login'});
    //     } else {
            next();
    //     }
    // } else {
    //     //  next({path: '/404'});
    // }
});

router.afterEach((to) => {

});
