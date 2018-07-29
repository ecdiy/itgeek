<template>
    <div>
        <Form :model="param" ref="form" :label-width="modParam.length>1? 120:10">
            <FormItem v-for="(it,pIdx) in modParam" :label="it.desc">
                <ele-html v-if="it.type=='html'" :v="param[it.name]" :paramName="it.name"
                          @on-result-change="onResultChange"></ele-html>
                <Input v-if="it.type=='text'" v-model="param[it.name]"></Input>

                <Select v-if="it.type=='select'" v-model="param[it.name]" style="width:200px">
                    <Option v-for="item in it.val" :value="item.val" :key="item.label">{{item.label}}</Option>
                </Select>
                <Card v-if="it.type=='array'">
                    <Row v-for="r in it.subParam">
                        <Col span="4">{{r.desc}}</Col>
                        <Col span="20">
                            <div style="width: 100%;padding:0 0 10px 0">
                                <!--<FormItem v-if="r.type=='text'" :label="r.name">-->
                                <!--<Input v-model="tmp[it.name][r.name]"></Input>-->
                                <!--</FormItem>-->
                                <ele-img v-if="r.type=='img'" :paramName="[it.name]"
                                         @on-result-change="onResultChange"></ele-img>
                            </div>
                            <Button style="margin-left: 60px" @click="arrayAction(it.name)">
                                添加
                            </Button>
                            <ul class="tag">
                                <li v-for="(item,index) in param[it.name]">
                                    <div v-if="item.url" class="hdf">
                                        <img :src="item.url" style="max-height: 100px;max-width: 100px"/>
                                        <div class="hd" title="删除" @click="handleDel(it.name,index)">
                                            <Icon type="close"></Icon>
                                        </div>
                                    </div>
                                    <div v-else>
                                        <!-- TODO  <div @click="selectTag(it.name,r.name,pIdx,index)" :class="cls[it.name][index]"-->
                                        <!--style="float: left;padding: 0 12px 0 8px;">-->
                                        <!--{{index+1}}-->
                                        <!--</div>-->
                                    </div>

                                </li>
                            </ul>
                        </Col>
                    </Row>
                </Card>
            </FormItem>
            <FormItem>
                <Button @click="save">保存</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>import eleImg from './ele-img';
import eleHtml from './ele-html';

export default {
    components: {eleImg, eleHtml},
    data() {
        return {
            param: {}, modParam: [], cls: {}, edit: {}, tmp: {}, editKey: ''
        }
    },
    created() {
        vm.$on("GetModuleParam", (ml, k, v) => {
            this.editKey = k;
            this.modParam = [];
            this.param = v;
            for (var i = 0; i < ml.length; i++) {
                if (k == ml[i].key) {
                    this.modParam = ml[i].param;
                    break;
                }
            }
            if (!this.param) {
                this.param = {}
            }
        })
    },
    methods: {
        handleDel(name, index) {
            this.param[name].splice(index, 1);
        },
        setImgData(n, v) {
            // if (n.length == 2) {
            //     this.tmp[n[0]][n[1]] = v;
            // } else {
            //     this.param[n[0]] = v;
            // }
        },
        arrayAction(pn) {
            vm.$emit("commitFormData")
            if (!this.param[pn]) {
                // this.param[pn] = [];
                this.$set(this.param, pn, [])
            }
            // if (this.modParam[pIdx].action == 'insert') {
            this.param[pn].push(JSON.parse(JSON.stringify(this.tmp[pn])));
            this.tmp[pn] = {};
            // } else {
            //     this.param[pn][pIdx] = this.tmp[pn];
            // }
            // this.$set(this.param, this.param);
        },
        onResultChange(k, v) {
            if (typeof(k) == 'object') {
                if (!this.tmp[k[0]]) {
                    this.tmp[k[0]] = {};
                }
                for (var kk in v) {
                    this.tmp[k[0]][kk] = v[kk]
                }
            } else {
                this.param[k] = v;
            }
        },
        save() {
            vm.$emit("commitFormData")
            vm.$emit("saveModuleParam", this.param)
        }
    }
}
</script>
<style scoped>
    ul.tag {
        clear: both;
        float: left;
        width: 100%;
    }

    ul.tag li {
        list-style: none;
        float: left;
        overflow: hidden;
        max-width: 120px;
        padding: 10px 3px;
    }

    .hd {
        display: none;
        position: relative;
        top: -35px;
        right: 0;
        background-color: #fff;
        width: 30px;
        height: 25px;
        float: right;
    }

    .hdf:hover .hd {
        display: block;
    }
</style>