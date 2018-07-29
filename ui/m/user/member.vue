<template>
    <div>
        <div class="box">
            <table cellpadding="0" cellspacing="0" border="0" width="100%">
                <tr>
                    <td width="50" align="left">
                        <img v-if="member.Id>0" :src="avatar(member.Id)" border="0">
                    </td>
                    <td>
                        <h4>{{member.Username}}</h4>
                    </td>
                </tr>
            </table>
            <div style="padding-left: 10px">
                加入于: {{member.CreateAt}}
                <router-link to="/p/user/dau">今日活跃度排名{{DauOrder}}</router-link>
            </div>
            <div>
                <x-button mini v-if="isLogin && followStatus>=0" @click="gz">{{gzAction}}</x-button>
            </div>
        </div>
        <div class="h10"/>
        <div class="box">
            <div class="h10"/>
            <topic-list :list="list" self="1"></topic-list>
        </div>
        <page v-if="totalTopicPage>1" :all="totalTopicPage" :no="current"
              :url="'/p/member/'+member.Username+','+tabs+',{}'"></page>

    </div>
</template>

<script>
    import topicList from '../topic/topicList'
    // import {ButtonTab, ButtonTabItem} from 'vux'

    export default {
        components: {
            topicList
            //     ButtonTab,
            //     ButtonTabItem
        }, computed: {
            totalTopicPage() {
                return Math.ceil(this.member.TopicCount / 20)
            }, gzAction() {
                return this.followStatus == 1 ? "取消关注" : (this.followStatus == 0 ? "关注" : "")
            },
            isLogin() {
                return window.gk.login && window.gk.user.Id != this.member.Id
            }
        },
        data() {
            return {
                followStatus: -1,
                member: {Id: 0, Username: '', CreateAt: '', Info: '', TopicCount: 0},
                tabs: 0, list: [], total: 0, current: 1, DauOrder: 0
            };
        },
        methods: {
            gz() {
                this.ajax('/gk-topic/follow', {userId: this.member.Id}, (o) => {
                    gk.user.FollowCount = o.followCount;
                })
            },
            load() {
                if (gk.login && gk.user.Username != this.member.Username) {
                    this.ajax('/gk-topic/followStatus', {id: this.member.Id})
                }
                this.loadData()
            },
            loadData() {
                if (this.tabs == 0) {
                    if (this.list.length == 0)
                        this.ajax('/gk-topic/listByUserId', {userId: this.member.Id, page: this.current})
                } else {
                    vm.$emit('memberNoteUserId', this.member.Id);
                }
            },
            loadTopic(data) {
                this.current = data;
                this.load()
            },
        },
        watch: {
            'tabs': 'loadData'
        },
        created() {
            var an = window.location.pathname.substr(10).split(',');
            if (an.length == 3) {
                this.current = Number(an[2])
            }
            this.ajax('/gk-user/memberInfo', {username: an[0]}, this.load)
        }
    }
</script>

<style scoped>

</style>