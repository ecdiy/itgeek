<template>
    <gk-body ec-key="TopicDetail">
        <div slot="left">
            <card>
                <row>
                    <Col span="18">
                        <Breadcrumb>
                            <BreadcrumbItem>
                                <router-link to="/">首页</router-link>
                            </BreadcrumbItem>
                            <BreadcrumbItem>
                                <router-link :to="'/p/topic/list,'+topic.ParentId+',0,1'">{{topic.ParentCatName}}
                                </router-link>
                            </BreadcrumbItem>
                            <BreadcrumbItem>
                                <router-link
                                        :to="'/p/topic/list,'+topic.ParentId+','+topic.CategoryId+',1'">
                                    {{topic.Name}}
                                </router-link>
                            </BreadcrumbItem>
                        </Breadcrumb>
                        <h2>{{topic.Title}}</h2>
                        <small class="gray">
                            <router-link :to="'/p/member/'+topic.Username">{{topic.Username}}</router-link>

                            <go :to="'topic/append,'+id" :title="'增加附言:'+topic.Title">APPEND</go>
                        </small>
                    </Col>
                    <Col span="6">
                        <router-link :to="'/p/member/'+topic.Username">
                            <img class="fr" :src="avatar(topic.UserId)" border="0"></router-link>
                    </Col>
                </row>
            </card>
            <card>
                <div class="markdown-body" v-html="topic.Body"></div>
            </card>

            <div v-for="(it,idx) in appendList" class="markdown-body">
                <card>
                    <span slot="title">第 {{idx+1}} 条附言  ·  {{it.CreateAt}}</span>
                    <div v-html="it.AppendText"></div>
                </card>
            </div>

            <div class="small" style="padding:5px 10px;border-radius: 4px;background-color:#e0e0e0; margin: 10px">
                    <span v-if="gk.login">
                        <a @click="doFav" v-if="favStatus==0">加入收藏</a>
                        <a @click="doFav" v-if="favStatus==1">取消收藏</a>
                    </span>
                <a :href="weiboUrl" target="_blank">微博分享</a>
                <span class="fr">{{topic.ShowTimes}} 次点击  {{topic.ReplyCount}} 回复</span>
            </div>
            <card v-for="(it,idx) in replyList">
                <table cellpadding="0" cellspacing="0" border="0" width="100%">
                    <tr>
                        <td width="48" valign="top" align="center">
                            <router-link :to="'/p/member/'+it.Username">
                                <img :src="avatar(it.UserId)" border="0"></router-link>

                        </td>
                        <td width="10"></td>
                        <td>
                            <div>
                                <router-link :to="'/p/member/'+it.Username">{{it.Username}}</router-link>
                                <span class="small">{{it.CreateAt}}</span>
                                <span class="fr">
                                    <span v-if="it.UserId!=gk.user.Id && !isThank(it.Id)">
                                    <!--<span class="hd">隐藏</span>-->
                                    <span class="hd"><a @click="doThank(it.Id)">感谢回复者</a></span>
                                    <Icon type="reply"></Icon>
                                    </span>
                                    <span class="no">{{idx+1}}</span>
                                </span></div>
                            <br>
                            <div v-html="it.Reply"></div>
                        </td>
                    </tr>
                </table>
            </card>

            <div v-if="gk.login" style="background:#eee;padding:10px 0">
                <card>
                    <div class="cell">
                        添加一条新回复
                    </div>
                </card>
                <card>
            <textarea name="content" maxlength="10000" v-model="fm.Reply"
                      style="overflow: hidden;width: 100%; word-wrap: break-word; resize: none; height: 112px;"></textarea>

                    <div class="fr">
                        <span class="gray">请尽量让自己的回复能够对别人有帮助</span></div>
                    <Button @click="postReply">回复</Button>
                </card>
            </div>
        </div>
        <div slot="right">
            <hot-topic></hot-topic>
        </div>
    </gk-body>
</template>

<script>
    import './md.css';
    import hotTopic from './hot'

    export default {
        components: {hotTopic},
        data() {
            return {
                thank: [], fm: {}, topic: {}, gk: window.gk, id: '', replyList: [], favStatus: -1,
                weiboUrl: 'https://service.weibo.com/share/share.php?url=' + encodeURIComponent(window.location),
                appendList: []
            }
        }, created() {
            vm.$on("data", (p) => {
                this.gk = p;
            })
        }, methods: {
            doFav() {
                this.ajax('/gk-topic/fav', {id: this.id}, function (r, th) {
                    th.favStatus = r.Fav
                    vm.$emit("TopicFavCount", r.TopicFavCount);
                })
            }, init() {
                var hash = window.location.pathname;
                var p = hash.split(",")
                this.id = p[1]
                this.ajax('/gk-topic/detail', {id: this.id, uId: p[2], referer: document.referrer}, (r, th) => {
                    th.topic = r.Topic;
                    th.replyList = r.Reply;
                    th.weiboUrl += '&title=' + encodeURIComponent(th.topic.Title);
                })
                if (window.gk.login) {
                    this.ajax('/gk-topic/favStatus', {id: this.id}, (r, th) => {
                        th.favStatus = r.Fav
                    })
                }
            }, postReply: function () {
                if (gk.user.Score <= 0) {
                    this.$Notice.error({
                        title: '积分不够'
                    });
                    return
                }
                this.fm.TopicId = this.id;
                this.ajax('/gk-topic/reply', this.fm, (r, th) => {
                    th.fm.Reply = "";
                    th.topic.ReplyCount = r.Count;
                    th.replyList = r.list;
                    gk.user.Score = r.Score;
                })
            },
            isThank(id) {
                if (this.thank) {
                    for (var i = 0; i < this.thank.length; i++) {
                        if (this.thank[i] == id) {
                            return true;
                        }
                    }
                }
                return false
            }, doThank(id) {
                this.ajax('/gk-topic/topicReplyThank', {"ReplyId": id}, (r, th) => {
                    if (r.score) {
                        gk.user.Score = r.score;
                    }
                });
            }
        }, mounted() {
            this.init();
        }, activated() {
            this.init();
        }, watch: {
            // 如果路由有变化，会再次执行该方法
            '$route': 'init'
        },
    }
</script>

<style>
    div.ivu-card-body {
        overflow: hidden;
    }

    .no {
        font-size: 9px;
        line-height: 9px;
        font-weight: 500;
        border-radius: 5px;
        display: inline-block;
        background-color: #f0f0f0;
        color: #aaa;
        padding: 5px 6px;
        cursor: pointer;
        margin-left: 10px;
    }

    .hd {
        display: none;
        float: left;
        padding-right: 10px;
    }

    .fr:hover .hd {
        display: block;
    }
</style>