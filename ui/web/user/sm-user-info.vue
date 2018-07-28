<template>
    <div>
        <card>
            <div style="clear: both;height: 48px">
                <div style="padding: 10px 0;line-height: 48px" class="fl">
                    <router-link :to="'/p/member/'+gk.user.Username">
                        <img :src="hdImg" border="0" width="48" height="48"/>
                    </router-link>
                </div>
                <div style="padding: 10px;" class="fl">
                    <div>
                        <router-link style="padding: 0;margin: 0" :to="'/p/member/'+gk.user.Username"><h2>
                            {{gk.user.Username}}</h2>
                        </router-link>
                    </div>
                    <div style="clear: both;padding: 0" class="small">
                        <router-link to="/p/my/balance">积分：{{gk.user.Score}}</router-link>
                    </div>
                </div>
            </div>

            <div style="clear: both;">
                <Row>
                    <Col span="12">
                        <router-link to="/p/my/fav">{{gk.user.FavCount}}<br>主题收藏</router-link>
                    </Col>
                    <Col span="12">
                        <router-link to="/p/my/follow">{{gk.user.FollowCount}}<br>特别关注</router-link>
                    </Col>
                </Row>
            </div>
        </card>
        <card>
            <a @click="openNew">创作新主题</a>
        </card>
        <card v-if="gk.user.LoginAward==1">
            <go to="my/LoginAward" title="领取今日的登录奖励"><Icon type="cube" size="18"></Icon> 领取今日的登录奖励</go>
        </card>
        <card>
            <router-link to="/p/my/msg" class="message-con">
                <span v-if="gk.user.MsgCount > 0"><Icon type="speakerphone" size="18"></Icon>有{{msgUnRead}} 条未读消息</span>
                <span v-else>无未读消息</span>
            </router-link>
        </card>
        <card>
            <go to="topic/referer">来源分析</go>
        </card>
    </div>
</template>
<script>
    export default {
        data() {
            return {gk: window.gk, v: 0}
        }, created() {
            vm.$on("data", (p) => {
                this.gk = p;
                this.v++;
            })

            vm.$on("TopicFavCount", (p) => {
                this.gk.user.FavCount = p;
            })
        }, computed: {
            hdImg() {
                return this.avatar(gk.user.Id) + this.v
            },
            msgUnRead() {
                return this.gk.user.MsgCount
            },
        }, methods: {
            openNew() {
                if (gk.user.Score <= 0) {
                    this.$Notice.error({
                        title: '积分不够'
                    });
                    return
                }
                this.$router.push('/p/new' + gk.user.EditType)
            },
        }
    }
</script>