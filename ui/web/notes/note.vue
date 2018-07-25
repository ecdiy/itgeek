<template>
    <row>
        <Col span="6">
            <div class="xtree">
                <!--<card>-->
                <!--<router-link :to="'/p/notes/new' +  user.EditType">-->
                <!--<Button>新建笔记</Button>-->
                <!--</router-link>-->
                <!--<router-link to="/p/notes/category">-->
                <!--<Button>分类管理</Button>-->
                <!--</router-link>-->
                <!--</card>-->
                <card>
                    <Tree :data="catTree" :render="renderContent"></Tree>
                </card>
            </div>
        </Col>
        <Col span="18">
            <card>
                <card>
                    <h2>{{notes.Title}}</h2>
                </card>
                <card>
                    <div v-html="notes.Body"></div>
                </card>
            </card>
        </Col>
    </row>
</template>
<script>
    export default {
        data() {
            return {catTree: [], notes: {}, userId: ''};
        }, methods: {
            treeCat: function (fId, r, nod) {
                for (var i = 0; i < r.cat.length; i++) {
                    var it = r.cat[i];
                    if (it.ParentId == fId) {
                        var o = {
                            title: it.Name, expand: true, children: [], type: 'cat'
                        };
                        this.treeCat(it.Id, r, o.children)
                        nod.push(o)
                    }
                }
                for (var i = 0; i < r.list.length; i++) {
                    var it = r.list[i];
                    if (it.CategoryId == fId) {
                        nod.push({title: it.Title, type: 'item', Id: it.Id})
                    }
                }
            }
            , renderContent(h, {root, node, data}) {
                return h('span', {
                    style: {
                        display: 'inline-block', width: '100%'
                    }
                }, [
                    h('span', [
                        h('Icon', {
                            props: {
                                type: data.type == 'cat' ? 'folder' : 'ios-paper-outline'
                            },
                            style: {
                                marginRight: '8px'
                            }
                        }),
                        h('span', {
                            on: {
                                click: (p) => {
                                    if (data.type == 'item') {
                                        this.ajax('/gk-note/user/detail', {
                                            id: data.Id,
                                            userId: this.userId
                                        }, (r, th) => {
                                            th.notes = r.Notes
                                        })
                                    }
                                }
                            }
                        }, data.title)
                    ]),
                    h('span', {
                        style: {
                            display: 'inline-block', float: 'right', marginRight: '32px'
                        }
                    })
                ]);
            },
        }, created() {
            vm.$on('memberNoteUserId', (id) => {
                if (this.userId == id) return;
                this.catTree = [];
                this.userId = id;
                if (id > 0) {
                    this.ajax('/gk-note/user/list', {"userId": id}, (r, th) => {
                        th.treeCat(0, r, th.catTree)
                    });
                } else {
                    this.catTree = [];
                    this.notes = {}
                }
            });
        }
    }
</script>
<style>
    .xtree ul.ivu-tree-children {
        clear: both;
    }

    .xtree .ivu-card-bordered {
        float: left;
        display: block;
        width: 100%;
    }
</style>