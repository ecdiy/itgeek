<template>
    <div style="background:#eee;padding: 10px 0;clear: both;position: relative">

        <Row>
            <Col span="6">
                <card>
                    <Sider hide-trigger :style="{background: '#fff'}">
                        <Menu :active-name="an" theme="light" width="auto" :open-names="on" @on-select="changeMenu">
                            <Submenu :name="idx" v-for="(it,idx) in menu">
                                <template slot="title">
                                    <Icon type="ios-analytics"></Icon>
                                    {{it.name}}
                                </template>
                                <MenuItem v-for="(it2,idx2) in it.sub" :name="idx+'-'+idx2">
                                    {{it2.name}}
                                </MenuItem>
                            </Submenu>
                        </Menu>
                    </Sider>
                </card>
            </Col>
            <Col span="18">
                <div style="background:#eee;padding:0 0 0 10px">
                    <card>
                        <keep-alive>
                            <router-view></router-view>
                        </keep-alive>
                    </card>
                </div>
            </Col>
        </Row>


    </div>

</template>

<script>
    export default {
        data() {
            return {
                an: '', on: '',
                gk: window.gk, menu: [
                    {
                        name: "社区管理", sub: [
                            {name: '积分规则', path: 'gkadmin/scoreRule'},
                            {name: '基础设置', path: 'gkadmin/index'},
                            {name: '分类管理', path: 'gkadmin/category'},
                            {name: '用户管理', path: 'gkadmin/user'},
                            // {name: '主题管理', path: 'gkadmin/topic'},
                        ]
                    },
                    {
                        name: "Web页面管理", sub: [
                            {name: '首页', path: 'gkadmin/pageIndex'},
                            {name: '主题详情', path: 'gkadmin/pageTopicDetail'},
                        ]
                    }, {
                        name: '账户', sub: [
                            {name: '修改密码', path: 'gkadmin/upUserPass'}
                        ]
                    }
                ]
            }
        }, created() {
            vm.$on("data", (p) => {
                this.gk = p;
            })
            var pn = window.location.pathname
            window.gk.siteId = pn.split(",")[1]
            for (var i = 0; i < this.menu.length; i++) {
                for (var j = 0; j < this.menu[i].sub.length; j++) {
                    if (pn.indexOf(this.menu[i].sub[j].path) > 0) {
                        this.an = i + "-" + j;
                        this.on = [i];
                    }
                }
            }
        }, methods: {
            changeMenu(p) {
                var i = p.indexOf("-");
                var px = this.menu[p.substr(0, i)].sub[p.substr(i + 1)]
                document.title = px.name;
                vm.$router.push({path: '/p/' + px.path})
            }
        }
    }
</script>