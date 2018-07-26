<template>
    <div>
        <card>
            <Breadcrumb>
                <BreadcrumbItem>
                    <router-link to="/">首页</router-link>
                </BreadcrumbItem>
                <!--<BreadcrumbItem>-->
                <!--<router-link :to="'/p/topic/list,'+topic.ParentId+',0,1'">{{topic.ParentCatName}}-->
                <!--</router-link>-->
                <!--</BreadcrumbItem>-->
                <BreadcrumbItem>
                    <router-link :to="'/p/topic/detail,' + topicId + ',' + userId">
                        {{title}}
                    </router-link>
                </BreadcrumbItem>
            </Breadcrumb>
        </card>
        <Card shadow>
            <!--<span slot=""></span>-->
            <textarea class='tinymce-textarea' id="tinymceEditer"></textarea>
            <Button @click="saveGoods">保存</Button>
            <Spin fix v-if="spinShow">
                <Icon type="load-c" size=18 class="demo-spin-icon-load"></Icon>
                <div>加载组件中...</div>
            </Spin>
        </Card>
        <Card>
            <span slot="title">关于为主题创建附言</span>
            请在确有必要的情况下再使用此功能为原主题补充信息<br>
            每个主题至多可以附加 3 条附言<br>
            创建附言价格为 20
        </Card>
    </div>
</template>

<script>
    import tinymce from 'tinymce';

    export default {
        data() {
            return {fm: {}, imgParam: {}, spinShow: true, topicId: '', userId: '', title: ''};
        },
        methods: {
            init() {
                var p = window.location.pathname.split(",")
                this.topicId = p[1];
                this.userId = p[2];
                var t = window.document.title;
                var idx = t.indexOf(":");
                if (idx > 0) {
                    this.title = t.substr(idx + 1)
                } else {
                    this.ajax('/gk-topic/topicBase', {Id: this.topicId}, (r, th) => {
                        th.title = r.base.Title;
                    })
                }
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
                this.fm.AppendText = html;
                this.fm.TopicId = this.topicId
                this.ajax("/gk-topic/topicAppend", this.fm, (d, th) => {
                    localStorage.editorContent = "";
                    gk.user.Score = d.score;
                    th.$router.replace('/p/topic/detail,' + th.topicId + ',' + th.userId)
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
