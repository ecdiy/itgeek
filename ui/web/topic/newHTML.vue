<template>
    <Row>
        <Col span="18">
            <Card shadow>

                <Form :label-width="80">
                    <FormItem label="标题">
                        <Input v-model="fm.Title"/>
                    </FormItem>
                    <FormItem label="选择分类">
                        <Cascader :data="catTreeList" v-model="catId"></Cascader>
                    </FormItem>
                </Form>
                <textarea class='tinymce-textarea' id="tinymceEditer"></textarea>


                <Button @click="saveGoods">保存</Button>
            </Card>
            <Spin fix v-if="spinShow">
                <Icon type="load-c" size=18 class="demo-spin-icon-load"></Icon>
                <div>加载组件中...</div>
            </Spin>
        </Col>
        <Col span="6">
            <div style="background:#eee;padding:0 0 0 10px">
                <card></card>
            </div>
        </Col>
    </Row>
</template>

<script>
    import tinymce from 'tinymce';

    export default {
        data() {
            return {fm: {}, imgParam: {}, spinShow: true, catTreeList: [], catId: []};
        },
        methods: {
            init() {
                this.ajax('/gk-topic/category/list', {}, (r,th) => {
                    for (var i = 0; i < r.catList.length; i++) {
                        var it = r.catList[i];
                        if (it.ParentId == 0) {
                            var o = {value: it.Id, label: it.Name, children: []};
                            for (var j = 0; j < r.catList.length; j++) {
                                var it2 = r.catList[j];
                                if (it2.ParentId == it.Id) {
                                    o.children.push({value: it2.Id, label: it2.Name})
                                }
                            }
                            th.catTreeList.push(o);
                        }
                    }
                });

                this.$nextTick(() => {
                    let vm = this;
                    let height = document.body.offsetHeight - 300;
                    tinymce.init({
                        selector: '#tinymceEditer', branding: false, elementpath: false,
                        height: height, language: 'zh_CN.GB2312',
                        menubar: 'edit insert view format table tools',
                        plugins: [
                            'advlist autolink lists link image charmap print preview hr anchor pagebreak imagetools',
                            'searchreplace visualblocks visualchars code fullpage',
                            'insertdatetime media nonbreaking save table contextmenu directionality',
                            'emoticons paste textcolor colorpicker textpattern imagetools codesample'
                        ],
                        toolbar1: ' newnote preview | undo redo | insert | styleselect | forecolor backcolor bold italic | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link image media',
                        autosave_interval: '20s',
                        image_advtab: true,
                        table_default_styles: {width: '100%', borderCollapse: 'collapse'},
                        images_upload_url: "/api/gk-upload/upload",
                        setup: function (editor) {
                            editor.on('init', function (e) {
                                vm.spinShow = false;
                                if (localStorage.editorContent) {
                                    tinymce.get('tinymceEditer').setContent(localStorage.editorContent);
                                }
                            });
                            editor.on('keydown', function (e) {
                                localStorage.editorContent = tinymce.get('tinymceEditer').getContent();
                            });
                        }
                    });
                });

            },

            saveGoods() {
                var html = tinymce.activeEditor.getContent();
                var idx = html.indexOf("<body>");
                if (idx > 0) {
                    html = html.substr(idx + 6)
                }
                idx = html.indexOf("</body>");
                if (idx > 0) {
                    html = html.substr(0, idx)
                }
                this.fm.SourceType = 'Html';
                this.fm.Body = html;
                if (this.catId.length == 0) {
                    return
                }
                this.fm.CategoryId = this.catId[1];
                this.ajax("/gk-topic/save", this.fm, function (d,th) {
                    localStorage.editorContent = "";
                    th .$router.replace('/')
                });
            },

            handleUploadImg(res, file, f) {
                this.url = file.response.uri;
                this.param.url = this.url;
                this.$emit('setImgData', this.paramName, this.param);
            },
        },
        mounted() {
            this.init();
        }, destroyed() {
            tinymce.get('tinymceEditer').destroy();
        }
    };
</script>
