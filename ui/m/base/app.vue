<template>
    <div id="main" class="app-main">
        <div class="cell hd">

             <span class="fl" v-if="!gk.site.Logo || gk.site.Logo==''" style="font-weight: bold">
                 <router-link to="/">{{gk.site.SiteName}}</router-link>
              </span>


            <router-link v-if="gk.site.Logo && gk.site.Logo!=''" class="fl" to="/" :title="gk.site.SiteName">
                <img :src="gk.site.Logo" height="50"/>
            </router-link>
            <span class="fr" style="display: block">
                 <span v-if="gk.login">
                        {{gk.user.Username}}
                     <!--<router-link to="#/notes/list">笔记</router-link>-->
                     <!--<router-link to="#/my/setting">设置</router-link>-->
                        <a @click="loginOut">登出</a>
                    </span>
                    <span v-else>
                        <router-link to="/">首页</router-link>
                        <router-link to="/p/user/register">注册</router-link>
                        <router-link to="/p/user/login">登录</router-link>
                    </span>
            </span>
        </div>
        <router-view></router-view>
    </div>
</template>

<script>import Cookies from 'js-cookie';

export default {
    data() {
        return {gk: window.gk}
    }, beforeCreate() {
        this.ajax('/gk-user/site', {}, (r, t) => {
            if (r.site && r.site.length > 2) {
                gk.site = JSON.parse(r.site);
                gk.siteId = r.SiteId;
                document.title = gk.site.SiteName;
            }
            if (r.Login) {
                gk.login = true;
                gk.user = r.Info;
            } else {
                Cookies.remove("h5Token");
                gk.user = {};
                gk.login = false;
            }
            vm.$emit("data", gk)
        })
    }, methods: {
        loginOut() {
            window.gk.login = false;
            Cookies.remove("username");
            Cookies.remove("h5Token");
            vm.$emit("login", false);
        }
    }
};
</script>
<style>
    a:active, a:link, a:visited {
        color: #778087;
        text-decoration: none;
        word-break: break-all;
    }

    .hd {
        height: 25px;
        line-height: 25px;
    }

    .app-main {
        word-break: break-all;
        width: 100%;
        height: 100%;
    }

    .nav {
        clear: both;
        float: left;
        display: block;
        padding: 10px;
    }

    .nav span {
        font-weight: 500;
        padding: 0 10px;
    }

    .bar {
        display: block;
        width: 100%;
        clear: both;
        float: left;
    }

    .cell {
        display: block;
        width: 100%;
        clear: both;
        float: left;
        padding: 10px 0;
        line-height: 120%;
        text-align: left;
        border-bottom: 1px solid #e2e2e2;
    }

    .fl {
        float: left;
    }

    .fr {
        padding-right: 10px;
        float: right;
    }

    li {
        list-style: none;
        /*float: left;*/
    }

    .small {
        font-size: 11px;
    }
</style>
