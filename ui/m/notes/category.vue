<template>
    <div>
        <div v-if="it.ParentId==0" v-for="(it,index) in list">


            <flexbox v-if="it.Id==editId">
                <flexbox-item>
                    <x-input v-model="editV"/>
                </flexbox-item>
                <flexbox-item>
                    <x-button mini @click.native="editSave">修改</x-button>
                    <x-button mini @click.native="cancelEdit">取消修改</x-button>
                </flexbox-item>
            </flexbox>

            <flexbox v-else>
                <flexbox-item>{{it.Name}}</flexbox-item>
                <flexbox-item>
                    <x-button mini @click.native="edit(it.Id,it.Name,index)">修改</x-button>
                    <x-button mini @click.native="catDel(it.Id,index)">删除</x-button>
                </flexbox-item>
            </flexbox>

            <div class="sub">
                <div class="bar" v-if="i.ParentId==it.Id" v-for="(i,idx) in list">
                    <flexbox v-if="i.Id==editId">
                        <flexbox-item>
                            <x-input mini placeholder="分类名称" v-model="editV"/>
                        </flexbox-item>
                        <flexbox-item>
                            <x-button mini @click.native="editSave">修改</x-button>
                            <x-button mini @click.native="cancelEdit">取消修改</x-button>
                        </flexbox-item>
                    </flexbox>
                    <flexbox v-else>
                        <flexbox-item>
                            {{i.Name}}
                        </flexbox-item>
                        <flexbox-item>
                            <x-button mini @click.native="edit(i.Id,i.Name,idx)">修改</x-button>
                            <x-button mini @click.native="catDel(i.Id,idx)">删除</x-button>
                        </flexbox-item>
                    </flexbox>
                </div>
                <div v-if="editId==-1">
                    <flexbox>
                        <flexbox-item>
                            <x-input placeholder="输入子分类名称" v-model="n['x'+it.Id]"/>
                        </flexbox-item>
                        <flexbox-item>
                            <x-button mini @click.native="addCategory(it.Id)">添加</x-button>
                        </flexbox-item>
                    </flexbox>
                </div>
            </div>
        </div>
        <div v-if="editId==-1">
            <flexbox>
                <flexbox-item>
                    <x-input placeholder="输入分类名称" v-model="n['x0']"></x-input>
                </flexbox-item>
                <flexbox-item>
                    <x-button mini @click.native="addCategory(0)">添加</x-Button>
                </flexbox-item>
            </flexbox>
        </div>
    </div>
</template>

<script>
    export default {

        data() {
            return {
                editV: "", list: [], n: {}, editId: -1, sortMap: {}, editIdx: -1
            };
        }, methods: {
            cancelEdit() {
                this.editId = -1;
            }, edit(id, v, idx) {
                this.editIdx = idx;
                this.editV = v;
                this.editId = id;
            }, editSave() {
                this.ajax('/gk-topic/luax/catgory/modify', {Id: this.editId, Name: this.editV}, function (r, th) {
                    th.editId = -1;
                    th.list[th.editIdx].Name = th.editV;
                })
            }, undo() {
                this.editId = -1;
            }, addCategory(pId) {
                this.ajax('/gk-topic/luax/catgory/insert', {
                    "ParentId": pId, "Name": this.n['x' + pId],
                    "CatType": "notes", "action": "insert"
                }, (res, th) => {
                    th.list.push({
                        Id: res.Id, Name: th.n['x' + pId], ParentId: pId
                    });
                    th.n['x' + pId] = "";
                })
            }, catDel(id, idx) {
                this.$vux.confirm.show({
                    title: '删除',
                    content: '分类下的笔记不会删除,笔记将转移到未分类下',
                    onConfirm: () => {
                        this.ajax('/gk-topic/luax/notes/categoryDel', {"id": id}, function (res,th) {
                            th.list.splice(idx, 1);
                        })
                    }
                });
            }
        }, beforeRouteEnter(to, from, next) {
            next((v) => {
                this.ajax('/gk-topic/luax/catgory/list', {type: 'notes'}, function (r) {
                    v.list = r.list;
                });
            })
        }
    }
</script>

<style scoped>

    .sub {
        padding: 0 25px;
        clear: both;
    }

    .sub div {
        clear: both;
    }
</style>