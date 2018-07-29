<template>

    <div style="width: 500px">
        <div v-if="it.ParentId==0" v-for="(it,index) in list"
             style="margin: 5px; padding: 10px">
            <span class="tit">{{it.Name}}</span>

            <span v-if="it.Id==editId">
                                <Input style="width: 100px" v-model="editV"/>
                                <Button @click="editSave">修改</Button>
                                <i title="取消修改" style="float: none;padding: 0 10px" @click="cancelEdit"><Icon
                                        type="android-cancel"></Icon></i>
             </span> <span v-else>
                        <span @click="edit(it.Id,it.Name,index)"><Icon type="edit"></Icon></span>
                        <span @click="catDel(it.Id,index)"><Icon type="android-remove-circle"></Icon></span>
                    </span>
            <div class="sub">
                <div v-if="i.ParentId==it.Id" v-for="(i,idx) in list">
                            <span v-if="i.Id==editId">
                                <Input style="width: 100px" v-model="editV"/>
                                <Button @click="editSave">修改</Button>
                                <i title="取消修改" style="float: none;padding: 0 10px" @click="cancelEdit"><Icon
                                        type="android-cancel"></Icon></i>
                            </span>
                    <span v-else>
                            <span class="tit">{{i.Name}}</span>
                            <span @click="edit(i.Id,i.Name,idx)"><Icon type="edit"></Icon></span>
                            <span @click="catDel(i.Id,idx)"><Icon type="android-remove-circle"></Icon></span>
                            </span>
                </div>
                <div v-if="editId==-1">
                    <Input style="width: 140px" v-model="n['x'+it.Id]">
                      <Button slot="append" @click="addCategory(it.Id)">+</Button>
                    </Input>
                </div>
            </div>
        </div>
        <div v-if="editId==-1">
            <Input style="width: 150px" v-model="n['x0']">
            <Button slot="append" @click="addCategory(0)">+</Button>
            </Input>
        </div>
    </div>

</template>

<script>
    export default {
        data() {
            return {
                editV: "", list: [], n: {}, editId: -1, sortMap: {}, editIdx: -1
            };
        }, created() {
            this.ajax('/gk-admin/topicCategoryList', {}, (r, th) => {
                th.list = r.catList;
            });
        }, methods: {
            cancelEdit() {
                this.editId = -1;
            }, edit(id, v, idx) {
                this.editIdx = idx;
                this.editV = v;
                this.editId = id;
            }, editSave() {
                this.ajax('/gk-admin/topicCategoryModify', {Id: this.editId, Name: this.editV}, (r, th) => {
                    th.editId = -1;
                    th.list[th.editIdx].Name = th.editV;
                })
            }, undo() {
                this.editId = -1;
            }, addCategory(pId) {
                this.ajax('/gk-admin/topicCategoryAdd', {
                    "ParentId": pId, "Name": this.n['x' + pId],

                }, (res, th) => {
                    th.list.push({
                        Id: res.Id, Name: th.n['x' + pId], ParentId: pId
                    });
                    th.n['x' + pId] = "";
                })
            }, catDel(id, idx) {
                let th = this;
                this.$Modal.confirm({
                    title: '删除', content: '<p>分类下的主题不会删除,将转移到未分类下</p>',
                    onOk() {
                        th.list.splice(idx, 1);
                        this.ajax('/gk-admin/topicCategoryDel', {"CategoryId": id}, (res, th) => {

                        })
                    }
                });
            }
        }
    }
</script>

<style scoped>
    #categoryList li {
        clear: both;
        margin: 5px;
        list-style: none;
        line-height: 32px;
    }

    div {
        padding: 5px 0 0;

    }

    span.tit {
        width: 160px;
        clear: both;
    }

    span {
        display: block;
        padding: 0 10px 0 0;
        float: left;
    }

    .sub {
        padding: 0 25px;
        clear: both;
    }

    .sub div {
        line-height: 36px;
        height: 36px;
    }

    .sub div {
        clear: both;
    }
</style>