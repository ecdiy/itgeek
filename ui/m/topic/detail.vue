<template>
    <div class="bar">
        <div class="nav">
            <router-link v-if="!gk.site.Logo || gk.site.Logo==''" class="fl" to="/">
               首页
            </router-link>

            <span>&nbsp;›&nbsp;</span>
            <router-link :to="'/p/topic/list,'+topic.ParentId+',0,1'">{{topic.ParentCatName}}</router-link>
            <span>&nbsp;›&nbsp;</span>
            <router-link :to="'/p/topic/list,0,'+topic.CategoryId+',1'">{{topic.Name}}</router-link>
        </div>

        <div class="cell">
            <h2>{{topic.Title}}</h2>
            <small class="gray">
                <router-link :to="'/p/member/'+topic.Username">{{topic.Username}}</router-link>
            </small>
        </div>
        <div class="cell markdown-body" v-html="topic.Body">
        </div>
        <div class="cell fav">
                    <span v-if="gk.login">
                        <a @click="doFav" v-if="favStatus==0">加入收藏</a>
                        <a @click="doFav" v-if="favStatus==1">取消收藏</a>
                    </span>
            <span class="fr">{{topic.ShowTimes}} 次点击  {{topic.ReplyCount}} 回复</span>
        </div>
        <div class="cell" v-for="it in replyList">
            <table cellpadding="0" cellspacing="0" border="0" width="100%">
                <tr>
                    <td width="48" valign="top" align="center">
                        <router-link :to="'/p/member/'+it.Username">
                            <img :src="avatar(it.UserId)" border="0"></router-link>
                    </td>
                    <td width="10"></td>
                    <td>
                        <router-link :to="'/p/member/'+it.Username">{{it.Username}}</router-link>
                        <span
                                class="fade small">{{ it.CreateAt}}</span><br>
                        <div v-html="it.Reply"></div>
                    </td>
                </tr>
            </table>
        </div>
        <div class="bar" v-if="gk.login">
            <x-textarea :placeholder="'添加一条新回复'" v-model="fm.Reply"></x-textarea>
            <x-button @click.native="postReply">回复</x-button>
        </div>
    </div>
</template>

<script> import {XTextarea} from 'vux'
import './md.css'

export default {
    components: {XTextarea},
    data() {
        return {topic: {}, gk: window.gk, id: '', replyList: [], favStatus: -1, fm: {}}
    }, created() {
        vm.$on("data", (p) => {
            this.gk = p;
        })
        this.init();
    }, watch: {
        // 如果路由有变化，会再次执行该方法
        '$route': 'init'
    }, methods: {
        postReply() {
            if (!this.fm.Reply && this.fm.Reply.length < 5) {
                return
            }
            this.fm.TopicId = this.id;
            this.ajax('/gk-topic/reply', this.fm, (r, th) => {
                th.replyList = r.list
                th.fm.Reply = "";
                th.topic.ReplyCount = r.Count
            })
        },
        doFav() {
            this.ajax('/gk-topic/fav', {id: this.id}, function (r, th) {
                th.favStatus = r.Fav
                vm.$emit("TopicFavCount", r.TopicFavCount);
            })
        }, init() {
            var hash = window.location.pathname;
            var p = hash.split(",")
            this.id = p[1]
            this.ajax('/gk-topic/detail', {id: this.id, uId: p[2]}, (r, th) => {
                // var hh = window.location.hash;
                // var idx = hh.indexOf(";");
                // if (idx > 0) {
                //     hh = hh.substr(0, idx)
                // }
                // hh += ";"
                // r.Topic.Body = r.Topic.Body.replaceAll('<router-link to="#',
                //     '<router-link to="' + hh).replaceAll('<router-link id="', '<router-link id="' + hh.substr(1));
                th.topic = r.Topic;
                th.replyList = r.Reply;
            })
            if (window.gk.login) {
                this.ajax('/gk-topic/favStatus', {id: this.id}, (r, th) => {
                    th.favStatus = r.Fav
                })
            }


        }
    }
}
</script>

<style>  td {
    vertical-align: top;
}

.fav, .fav a {
    font-size: 12px;
}
</style>