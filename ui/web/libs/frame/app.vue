<template>
    <div id="main" class="app-main">
        <card style="height: 50px;">
            <div class="content">
                <router-link v-if="!gk.site.Logo || gk.site.Logo==''" class="fl" to="/">
                    <h2>{{gk.site.SiteName}}</h2>
                </router-link>

                <router-link v-if="gk.site.Logo && gk.site.Logo!=''" class="fl" to="/" :title="gk.site.SiteName">
                    <img :src="gk.site.Logo" height="50"/>
                </router-link>


                <div class="h">
                    <ul>
                        <li>
                            <router-link to="/p/topic/list,0,0,1">主题</router-link>
                        </li>
                        <!--<li>文章</li>-->
                        <!--<li>-->
                        <!--<router-link to="/p/DevOps">项目/DevOps</router-link>-->
                        <!--</li>-->
                        <li>
                            <go to="note">热门笔记</go>
                        </li>
                        <li>
                            <go to="down">下载</go>
                        </li>
                    </ul>
                </div>

                <div class="fr">
                    <router-link to="/">首页</router-link>
                    <span v-if="gk.login">
                        {{gk.user.Username}}
                        <go to="notes/list">笔记</go>
                        <go to="my/setting">设置</go>
                          <a @click="loginOut">登出</a>
                    </span>
                    <span v-else>
                        <router-link to="/p/user/register">注册</router-link>
                        <router-link to="/p/user/login">登录</router-link>
                    </span>
                </div>
            </div>
        </card>
        <div class="content">
            <router-view></router-view>
        </div>
        <div class="content">
            <span>{{gk.site.Miibeian}}</span>
            <!--冀ICP备16008655号-3-->
        </div>
        <BackTop :height="100" :bottom="200">
            <div class="top">返回顶端</div>
        </BackTop>
    </div>
</template>

<script>
    import Cookies from 'js-cookie';

    export default {
        data() {
            return {gk: window.gk}
        }, methods: {
            loginOut() {
                gk.login = false;
                Cookies.remove("webToken");
                vm.$emit("login", false);
                this.$router.push("/")
            }
        }, beforeCreate() {
            this.ajax('/gk-user/site', {}, (r, t) => {
                if (r.site && r.site.length > 2) {
                    gk.site = JSON.parse(r.site);
                    gk.siteId = r.SiteId;
                    document.title = gk.site.SiteName;
                    if (!gk.site.dw) {
                        gk.site.dw = '铜币';
                    }
                }
                if (r.Login) {
                    gk.login = true;
                    gk.user = r.Info;
                } else {
                    Cookies.remove("webToken");
                    gk.user = {};
                    gk.login = false;
                }
            })
        }
    };
</script>

<style>
    html, body {
        width: 100%;
        height: 100%;
        background: #f0f0f0;
    }

    a:active, a:link, a:visited {
        color: #778087;
        text-decoration: none;
        word-break: break-all;
    }

    p {
        padding-top: 5px;
    }

    .app-main {
        width: 100%;
        height: 100%;
    }

    .fl {
        float: left;
    }

    .fr {
        float: right;
    }

    .small, .small a {
        padding-left: 10px;
        font-size: 11px;
        color: #646464;
        cursor: pointer;
    }

    .ft {
        padding: 10px 0;
    }

    .ft a {
        padding-left: 10px;
    }

    a {
        padding-left: 10px;
    }

    .content {
        min-width: 600px;
        max-width: 1100px;
        margin: 0 auto 0 auto;
        overflow: hidden;
    }

    .h li {
        float: left;
        padding-left: 40px;
        list-style: none;
        font-size: 18px;
    }

    .current {
        display: inline-block;
        font-size: 14px;
        line-height: 14px;
        padding: 5px 8px 5px 8px;
        margin-right: 5px;
        border-radius: 3px;
        background-color: #334;
        color: #fff;
    }

    a.current {
        color: #fff;
    }
</style>
