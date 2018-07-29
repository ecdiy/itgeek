<template>
    <div>
        <div class="box" v-if="gk.user.LoginAward==1">
            <div style="font-size: 14px;padding: 5px 10px 0;height: 18px;line-height: 18px">
                <x-icon type="cube" size="16"></x-icon>
                <go to="my/LoginAward">
                    领取今日的登录奖励
                </go>
            </div>
        </div>

        <div class="h10"/>

        <div class="box">
            <div class="cat cell">
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
                <ul>
                    <li v-if="it.ParentId!=0" v-for="it in catList">
                        <router-link v-if="pId==it.ParentId" :class="cId==it.Id?'current':''"
                                     @click="setId(it.ParentId,it.Id)"
                                     :to="'/p/topic/list,'+it.ParentId+','+it.Id+',1'">{{it.Name}}
                        </router-link>
                    </li>
                </ul>
            </div>
            <topic-list :list="topicList"></topic-list>
            <div class="cell">
                <page :all="totalPage" :no="current" :url="'/p/topic/list,'+pId+','+cId+',{}'"></page>
            </div>
        </div>
    </div>
</template>


<script>
    import topicList from '../topic/topicList'

    export default {
        components: {
            topicList
        },
        data() {
            return {topicList: [], pId: "0", cId: "0", catList: [], gk: window.gk, current: 1}
        },
        created() {
            this.cat();
            this.init();
            vm.$on("data", (p) => {
                this.gk = p;
            })
        },
        computed: {
            totalPage() {
                var t = 0;
                for (var i = 0; i < this.catList.length; i++) {
                    if ((this.cId > 0 && this.catList[i].Id == this.cId) || (this.cId == 0 && this.pId == 0) || (this.cId == 0 && this.pId == this.catList[i].ParentId)) {
                        t += Number(this.catList[i].ItemCount);
                    }
                }
                return Math.ceil(t / 20)
            }
        },
        watch: {
            '$route': 'init'
        }, methods: {
            cat() {
                if (this.catList.length == 0) {
                    this.ajax('/gk-topic/category/list')
                }
            },
            init() {
                var p = window.location.pathname.split(",");
                if (p.length > 3) {
                    this.pId = p[1];
                    this.cId = p[2];
                    this.current = Number(p[3]);
                }
                this.load()
            },
            upPid(v) {
                if (v.cId > 0 && v.pId == 0) {
                    for (var i = 0; i < v.catList.length; i++) {
                        if (v.catList[i].Id == v.cId) {
                            v.pId = v.catList[i].ParentId
                        }
                    }
                }
            },
            setId: function (p, r) {
                this.pId = p;
                this.cId = r;
                this.load()
            },
            load() {
                this.ajax('/gk-topic/list', {pId: this.pId, cId: this.cId, page: this.current}, (r, th) => {
                    th.upPid(th);
                })
            }
        }
    }
</script>
<style scoped>

    ul {
        padding: 0;
        margin: 0
    }

    .current {
        display: inline-block;
        border-radius: 3px;
        background-color: #334;
        color: #fff;
    }

    td {
        vertical-align: top;
    }

    td div {
        padding-top: 3px;
    }

    .cat {
        width: 100%;
        float: left;
        clear: both;
        display: block;
    }

    .cat li {
        display: block;
        display: block;
        margin: 0 5px;
        float: left;
    }

    li a {
        padding: 3px 4px 5px 4px;
    }

    .cat ul {
        float: left;
        display: block;
        clear: both;
        line-height: 120%;
        width: 100%;
    }

    table {
        display: table;
        border-collapse: separate;
        border-spacing: 2px;
        border-color: grey;
    }

    .fade {
        color: #ccc;
    }
</style>