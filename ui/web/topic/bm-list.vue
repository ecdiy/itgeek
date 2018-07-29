<template>
    <div>
        <card>
            <ul>
                <li>
                    <router-link :class="pId=='0' && cId=='0'?'current':''" @click="setId('0','0')"
                                 :to="'/p/topic/list,0,0,1'">所有
                    </router-link>
                </li>
                <li v-if="it.ParentId==0" v-for="it in catList">
                    <router-link :class="pId==it.Id?'current':''" @click="setId(it.Id,'0')"
                                 :to="'/p/topic/list,'+it.Id+',0,1'">{{it.Name}}
                    </router-link>
                </li>
            </ul>
            <ul v-if="pId>0">
                <li v-if="it.ParentId!=0" v-for="it in catList">
                    <router-link v-if="pId==it.ParentId" :class="cId==it.Id?'current':''"
                                 @click="setId(it.ParentId,it.Id)"
                                 :to="'/p/topic/list,'+it.ParentId+','+it.Id+',1'">{{it.Name}}
                    </router-link>
                </li>
            </ul>
        </card>
        <Card>
            <topic-list :list="topicList"></topic-list>
            <div style="padding: 10px 0">
                <Page :total="total" :page-size="20" @on-change="loadPage" :current="current"></Page>
            </div>
        </Card>

    </div>
</template>

<script>
    import topicList from '../topic/topicList'

    export default {
        components: {topicList},
        data() {
            return {topicList: [], pId: "0", cId: "0", catList: [],   current: 1}
        }, computed: {
            total () {
                var t = 0;
                for (var i = 0; i < this.catList.length; i++) {
                    if ((this.cId > 0 && this.catList[i].Id == this.cId) || (this.cId == 0 && this.pId == 0) || (this.cId == 0 && this.pId == this.catList[i].ParentId)) {
                        t += Number(this.catList[i].ItemCount);
                    }
                }
                return t
            }
        },
        methods: {
            loadPage(data) {
                // this.current = data;
                // this.load();
                vm.$router.replace('/p/topic/list,' + this.pId + ',' + this.cId + ',' + data)
            },
            cat() {
                if (this.catList.length == 0) {
                    this.ajax('/gk-topic/category/list')
                }
            },
            upPid() {
                var v = this;
                if (v.cId > 0 && v.pId == 0) {
                    for (var i = 0; i < v.catList.length; i++) {
                        if (v.catList[i].Id == v.cId) {
                            v.pId = v.catList[i].ParentId
                        }
                    }
                }
            },
            setId(p, r) {
                this.current = 1;
                this.pId = p;
                this.cId = r;
                this.load()
            },
            load() {
                this.ajax('/gk-topic/list', {pId: this.pId, cId: this.cId, page: this.current}, this.upPid)
            },
            init() {
                var pn = window.location.pathname;
                var p = pn.split(",")
                if (p.length > 3) {
                    this.pId = p[1];
                    this.cId = p[2];
                    this.current = Number(p[3]);
                }
                this.load()
            }
        },
        created() {
            this.cat();
            this.init()
        },
        watch: {
            '$route': 'init'
        }
    }
</script>

<style lang="less" scoped>
    ul {
        clear: both;
        height: 25px;
    }

    li {
        display: block;
        margin: 0 10px;
        float: left;
    }
</style>