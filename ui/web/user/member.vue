<template>
    <div>
        <Card>
            <card>
                <row>

                    <Col span="2" v-if="member.Id>0">
                        <img :src="avatar(member.Id)" border="0"></Col>
                    <Col span="20"><h2>{{member.Username}}</h2><br>
                        加入于: {{member.CreateAt}} 今日活跃度排名
                        <router-link to="/p/user/dau">{{DauOrder}}</router-link>
                    </Col>
                    <Col>
                        <Button v-if="isLogin && followStatus>=0" @click="gz">{{gzAction}}</Button>
                    </Col>

                </row>
            </card>
            <card>
                <member-info :json="member.Info"/>
            </card>

            <Tabs :animated="false" v-model="tabs">
                <TabPane label="创建所有主题">
                    <Card>

                        <topic-list :list="list" self="1"></topic-list>

                        <div style="padding: 10px 0">
                            <Page v-if="total>20" :total="total" :page-size="20" @on-change="loadTopic"
                                  :current="current"></Page>
                        </div>
                    </Card>
                </TabPane>
                <TabPane label="笔记">
                    <Card>
                        <note></note>
                    </Card>
                </TabPane>
            </Tabs>
        </Card>
    </div>
</template>

<script>
    import note from '../notes/note';
    import memberInfo from './_memberInfo'
    import topicList from '../topic/topicList'

    export default {
        components: {note, topicList, memberInfo},
        data() {
            return {
                followStatus: -1,
                member: {Id: 0, Username: '', CreateAt: '', Info: ''},
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
            var an = window.location.pathname.substr(10).split(',');
            if (an.length == 3) {
                this.current = Number(an[2])
            }
            this.ajax('/gk-user/memberInfo', {username: an[0]}, this.load)
        }
    }
</script>

<style lang="less" scoped>
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

    ul {
        clear: both;
        height: 25px;
    }

    li {
        display: block;
        margin: 0 10px;
    }

    .row {
        padding: 10px;
        font-size: 14px;
        line-height: 120%;
        text-align: left;
        border-bottom: 1px solid #e2e2e2;
    }

    .item_title {
        font-size: 16px;
        line-height: 130%;
        text-shadow: 0 1px 0 #fff;
        word-wrap: break-word;
        hyphens: auto;
        display: block;
        padding: 5px;
    }

    .topic_info {
        font-size: 12px;
        color: #ccc;
        line-height: 200%;
    }
</style>