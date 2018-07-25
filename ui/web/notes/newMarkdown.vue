<template>
    <div>
        <Row>
            <Col span="18">
                <Card shadow>
                    <Form :label-width="80">
                        <FormItem label="标题">
                            <Input v-model="fm.Title"/>
                        </FormItem>
                        <FormItem label="选择分类">
                            <Cascader :data="catList" v-model="catId"></Cascader>
                        </FormItem>
                    </Form>

                    <mavon-editor :defaultOpen="md.defaultOpen" ref="md" @imgAdd="$imgAdd" @save="save" v-model="value"
                                  :subfield="false"/>
                </Card>
            </Col>
            <Col span="6">
                <div style="background:#eee;padding:0 0 0 10px">
                    <card></card>
                </div>
            </Col>
        </Row>
    </div>
</template>

<script>
    import axios from "axios";
    import Vue from 'vue';
    import mavonEditor from 'mavon-editor'
    import 'mavon-editor/dist/css/index.css'

    Vue.use(mavonEditor)
    export default {
        data() {
            return {fm: {}, imgParam: {}, spinShow: true, catList: [], catId: [], value: '', md: {defaultOpen: "edit"}};
        },
        methods: {
            save(s, h) {
                this.md.defaultOpen = "preview";
                if (this.catId.length == 0 || this.fm.Title == "" || s == "") {
                    return
                }
                this.fm.SourceType = 'Markdown';
                this.fm.Body = h;
                this.fm.Source = s;
                this.fm.CategoryId = this.catId[1];
                this.ajax("/gk-note/add", this.fm, function (d, th) {
                    localStorage.editorContent = "";
                    th.$router.replace('/p/notes/list')
                });
            },
            $imgAdd(pos, $file) {
                let th = this;
                // 第一步.将图片上传到服务器.
                var formdata = new FormData();
                formdata.append('image', $file);
                axios({
                    url: '/gk-upload/upload', method: 'post', data: formdata,
                    headers: {'Content-Type': 'multipart/form-data'},
                }).then((url) => {
                    // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
                    // $vm.$img2Url 详情见本页末尾
                    th.$refs.md.$img2Url(pos, url.data.location);
                })
            }, init() {
                this.ajax('/gk-note/categoryList', {}, (r, th) => {
                    for (var i = 0; i < r.cat.length; i++) {
                        var it = r.cat[i];
                        if (it.ParentId == 0) {
                            var o = {value: it.Id, label: it.Name, children: []};
                            for (var j = 0; j < r.cat.length; j++) {
                                var it2 = r.cat[j];
                                if (it2.ParentId == it.Id) {
                                    o.children.push({value: it2.Id, label: it2.Name})
                                }
                            }
                            th.catList.push(o);
                        }
                    }
                });
            }
        }, mounted() {
            this.init();
        },
    }
</script>

<style>
    div.ivu-select-dropdown {
        z-index: 10000
    }

</style>