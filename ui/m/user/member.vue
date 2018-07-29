<template>
    <div>
        <div class="box">
            <div v-if="member.Id>0">
                <img :src="avatar(member.Id)" border="0"></div>
            <div><h2>{{member.Username}}</h2><br>
                加入于: {{member.CreateAt}} 今日活跃度排名
                <router-link to="/p/user/dau">{{DauOrder}}</router-link>
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
        },
        data() {
            return {
                followStatus: -1,
                member: {Id: 0, Username: '', CreateAt: '', Info: ''},
                tabs: 0, list: [], total: 0, current: 1, DauOrder: 0
            };
        },
        methods: {
            consoleIndex() {
            },
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
        }, computed: {
            gzAction() {
                return this.followStatus == 1 ? "取消关注" : (this.followStatus == 0 ? "关注" : "")
            },
            isLogin() {
                return window.gk.login && window.gk.user.Id != this.member.Id
            },
            TopicCount() {
                return Number(this.member.TopicCount);
            }
        },
        created() {
            var an = window.location.pathname.substr(10);
            this.ajax('/gk-user/memberInfo', {username: an}, this.load)
        }
    }
</script>

<style scoped>

</style>