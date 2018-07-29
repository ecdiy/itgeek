<template>
    <div>
        <card>
            <div class="m">
                <a @click="mLoad('all')" :class="mType=='all'?'current':''">所有({{totalAll}})</a>
                <a @click="mLoad('unRead')" :class="mType=='unRead'?'current':''">未读({{totalUnread}})</a>
            </div>
        </card>
        <card>
            <Row v-for="(it,idx) in msgList">
                <Col span="2">
                    <router-link :to="'/p/member/'+it.Username">
                        <img :src="avatar(it.FromUserId)" border="0"></router-link>
                </Col>
                <Col span="22" style="padding-left: 10px">
                    <div>
                        <span v-html="it.Title" :msg-id="it.Id" :class="it.Status=='0'?'unread':''"></span>
                        {{it.CreateAt}}
                        <Icon type="android-delete" @click="del(it.Id,idx)" size="16"></Icon>
                    </div>
                    <div class="payload" v-html="it.Body"></div>
                </Col>
            </Row>
            <Page v-if="total>20" :total="total" :page-size="20" :current="current" show-total @on-change="load"></Page>
        </card>
    </div>
</template>

<script>

    export default {
        data() {
            return {
                msgList: [], page: 1, current: 1, mType: 'all',
                total: 0, totalAll: 0, totalUnread: 0,
            }
        },
        created() {
            this.load(1)
            vm.$on("msgClick", (mId) => {
                for (var i = 0; i < this.msgList.length; i++) {
                    if (this.msgList[i].Id == mId) {
                        if (this.msgList[i].Status == '0') {
                            this.msgList[i].Status = 1;
                            this.ajax('/gk-user/msgRead', {id: mId})
                        }
                    }
                }
            })
        }, methods: {
            mLoad(m) {
                this.mType = m;
                this.load(this.current)
            },
            load(p) {
                this.current = p;
                this.ajax('/gk-user/msgList', {page: this.current, mType: this.mType})
            }, del(id, idx) {
                this.msgList.splice(idx, 1)
                this.ajax('/gk-user/msgDel', {id: id})
                if (this.total > 20) {
                    this.total = this.total - 1;
                }
            }
        }
    }
</script>
<style scoped>
    .m a {
        padding-left: 10px
    }

    .unread {
        font-weight: bold;
    }

    div.ivu-row {
        border-bottom: 1px solid #e2e2e2;
        padding: 8px;
    }

    .payload {
        display: inline-block;
        background-color: #f5f5f5;
        padding: 5px 10px 5px 10px;
        font-size: 14px;
        line-height: 120%;
        -moz-border-radius: 3px;
        -webkit-border-radius: 3px;
        border-radius: 3px;
        word-break: break-all;
        word-wrap: break-word;
    }
</style>