<template>
    <div id="main" class="app-main">
        <card style="height: 50px;">
            <div class="content">
                <go title="首页" class="fl" to="/">
                    <h3>后台管理</h3>
                </go>

                <div class="fr">
                    <span v-if="gk.login">
                        <go to="admin">后台管理</go>
                        <a @click="loginOut">登出</a>
                    </span>
                </div>
            </div>
        </card>
        <div class="content">
            <router-view></router-view>
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
                vm.$emit("data", window.gk)
                Cookies.remove("webGeekAdmin");
                vm.$router.replace('/p/login')
            }
        }, created() {
            this.ajax('/gk-admin/userStatus', {}, (r, th) => {
                if (!r.login) {
                    th.$router.replace('/p/login')
                    gk.login = false;
                } else {
                    th.gk.login = true;
                    gk.login = true;
                }
            });
            this.$on("data", (p) => {
                this.gk = p;
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
        margin-top: 10px;
        padding: 10px 0;
        background-color: #222;
        width: 100%;
        height: 140px;
        color: #fff;
        line-height: 30px;
    }

    .ft a {
        padding-left: 10px;
        color: #fff;
        line-height: 30px;
    }

    a {
        padding-left: 10px;
    }

    .content {
        min-width: 600px;
        max-width: 1100px;
        margin: 0 auto 0 auto;
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
