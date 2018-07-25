<template>
    <Row>
        <Col span="18">
            <card>
                <div v-if="it.ParentId==0" v-for="(it,index) in list">
                    <span class="tit">{{it.Name}}</span>

                    <span v-if="it.Id==editId">
                                <Input style="width: 100px" v-model="editV"/>
                                <Button @click="editSave">修改</Button>
                                <i title="取消修改" style="float: none;padding: 0 10px" @click="cancelEdit"><Icon
                                        type="android-cancel"></Icon></i>
                    </span>
                    <span v-else>
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
                            <Input style="width: 100px" v-model="n['x'+it.Id]"/>
                            <Button @click="addCategory(it.Id)">+</Button>
                        </div>
                    </div>
                </div>
                <div v-if="editId==-1">
                    <Input style="width: 100px" v-model="n['x0']"/>
                    <Button @click="addCategory(0)">+</Button>
                </div>
            </card>
        </Col>
        <Col span="6">
            <div style="background:#eee;padding:0 0 0 10px">
                <user-info></user-info>
            </div>
        </Col>
    </Row>
</template>

<script>
    export default {
        data() {
            return {
                editV: "", list: [], n: {}, editId: -1, sortMap: {}, editIdx: -1
            };
        }, created() {
            this.ajax('/gk-note/categoryList', {}, function (r, th) {
                th.list = r.cat;
            });
        }, methods: {
            cancelEdit() {
                this.editId = -1;
            }, edit(id, v, idx) {
                this.editIdx = idx;
                this.editV = v;
                this.editId = id;
            }, editSave() {
                this.ajax('/gk-note/categoryModify', {Id: this.editId, Name: this.editV}, function (r, th) {
                    th.editId = -1;
                    th.list[th.editIdx].Name = th.editV;
                })
            }, undo() {
                this.editId = -1;
            }, addCategory(pId) {
                this.ajax('/gk-note/categoryAdd', {
                    "ParentId": pId, "Name": this.n['x' + pId],

                }, function (res, th) {
                    th.list.push({
                        Id: res.Id, Name: th.n['x' + pId], ParentId: pId
                    });
                    th.n['x' + pId] = "";
                })
            }, catDel(id, idx) {
                this.$Modal.confirm({
                    title: '删除', content: '<p>分类下的笔记不会删除,笔记将转移到未分类下</p>',
                    onOk: () => {
                        this.ajax('/gk-note/categoryDel', {"id": id}, function (res, th) {
                            th.list.splice(idx, 1);
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
        clear: both;
    }
</style>