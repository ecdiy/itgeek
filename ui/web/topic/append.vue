<template>
    <div>
        <Card shadow>
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
            创建附言价格为每千字 20
        </Card>
    </div>
</template>

<script>
    import tinymce from 'tinymce';

    export default {
        data() {
            return {fm: {}, imgParam: {}, spinShow: true, Id: ''};
        },
        methods: {
            init() {
                this.Id = window.location.pathname.split[1];

                this.ajax('/gk-topic/topicBase', {id: this.id}, (r, th) => {
                    th.topic = r.Topic;
                    th.replyList = r.Reply;
                    th.weiboUrl += '&title=' + encodeURIComponent(th.topic.Title);
                })

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
                this.fm.Id = this.Id
                this.ajax("/gk-topic/append", this.fm, function (d, th) {
                    localStorage.editorContent = "";
                    th.$router.replace('/')
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
