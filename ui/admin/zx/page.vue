<template>
    <Card>
        <span slot="title">{{pageName}}</span>
        <div slot="extra"><a v-if="zxChange" @click.prevent="zxSave">
            <Icon type="checkmark-round"></Icon>
            保存</a>
            <a v-if="zxChange || zxPublic" @click.prevent="zxPub">
                <Icon type="ios-loop-strong"></Icon>
                发布</a>
        </div>
        <Row>
            <Col span="18">
                <card>
                    <div v-if="zx.main && zx.main.list">
                        <card v-for="(it,idx) in zx.main.list">
                        <span>{{it.name}}
                            <a @click="openParam('main',it,idx)">设置</a><a @click="delModule('main',idx)">删除</a>
                        </span>
                        </card>
                    </div>
                    <Button @click="openAddModal('main')">添加模块</Button>
                </card>
            </Col>
            <Col span="6">
                <card>
                    <div v-if="zx.right && zx.right.list">
                        <card v-for="(it2,idx2) in zx.right.list">
                        <span>{{it2.name}}
                            <a @click="openParam('right',it2,idx2)">设置</a><a @click="delModule('right',idx2)">删除</a>
                        </span>
                        </card>
                    </div>
                    <Button @click="openAddModal('right')">添加模块</Button>
                </card>
            </Col>
        </Row>
        <Modal v-model="  modulesModal" title="模块" :mask-closable="false">
            <div>
                <Button v-for="(it,v) in moduleList " @click="postModule(it,v)">{{it.name}}</Button>
            </div>
            <div slot="footer"></div>
        </Modal>
        <Modal v-model=" paramModal" :title="paramTitle+' - 模块参数'" :width="900" :mask-closable="false">
            <m-param></m-param>
            <div slot="footer"></div>
        </Modal>
    </Card>
</template>

<script>
    import mParam from "./moduleParam"
    import ModuleList from "./module.json";

    export default {
        components: {mParam},
        props: ['pageName', 'PageUa', 'PageKey'],
        data() {
            return {
                modulesModal: false, openKey: "", action: "",
                moduleList: ModuleList, zx: {right: {list: []}},
                zxPublic: true, zxChange: true, paramModal: false, editIndex: 0, paramTitle: ''
            }
        }, created() {
            this.ajax('/gk-admin/zxGet', {PageUa: this.PageUa, PageKey: this.PageKey}, (r, t) => {
                if (r.zxJson && r.zxJson.length > 2) {
                    t.zx = JSON.parse(r.zxJson)
                } else {
                    t.zx = {}
                }
            });
            vm.$on("saveModuleParam", (d) => {
                this.paramModal = false;
                this.zx[this.openKey].list[this.editIndex].val = d
                this.saveData()
            })
        }, methods: {
            openAddModal(k) {
                this.action = "add";
                this.openKey = k;
                this.modulesModal = true;
                if (!this.zx[this.openKey]) {
                    this.zx[this.openKey] = {list: []}
                }
            }, postModule(it) {
                this.modulesModal = false;
                this.zx[this.openKey].list.push({name: it.name, key: it.key, val: {}})
                this.saveData()
                var nod = this.zx[this.openKey].list.length - 1
                this.openParam(this.openKey, this.zx[this.openKey].list[nod], nod)
            },
            zxSave() {
                this.ajax('/gk-admin/zxSave', {
                    PageUa: this.PageUa, PageKey: this.PageKey,
                    Json: JSON.stringify(this.zx)
                })
                this.zxChange = false;
                this.zxPublic = true;
            },
            saveData() {
                this.zxChange = true;
            },
            zxPub() {
                this.ajax('/gk-admin/zxPub', {
                    PageUa: this.PageUa, PageKey: this.PageKey,
                })
                this.zxPublic = false;
            },
            delModule(key, idx) {
                this.zx[key].list.splice(idx, 1);
                this.zxChange = true;
                this.saveData();
            },
            openParam(k, it, i) {
                this.openKey = k;
                this.editIndex = i;
                this.paramModal = true;
                this.paramTitle = it.name;
                vm.$emit("GetModuleParam", ModuleList, it.key, JSON.parse(JSON.stringify(this.zx[this.openKey].list[this.editIndex].val)))
            }
        }
    }
</script>
