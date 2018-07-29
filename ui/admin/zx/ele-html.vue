<template>
    <div>
        <Card shadow>
            <textarea class='tinymce-textarea' id="tinymceEditer"></textarea>

        </Card>
    </div>
</template>

<script>import tinymce from 'tinymce';

export default {
    props: ['paramName', 'v'],
    watch: {
        v(vx) {
            tinymce.get('tinymceEditer').setContent(vx);
        }
    },
    created() {
        vm.$on("commitFormData", () => {
            var html = tinymce.activeEditor.getContent();
            var idx = html.indexOf("<body>");
            if (idx > 0) {
                html = html.substr(idx + 6)
            }
            idx = html.indexOf("</body>");
            if (idx > 0) {
                html = html.substr(0, idx)
            }
            this.$emit("on-result-change", this.paramName, html);
        });
        let th = this;
        this.$nextTick(() => {
            let height = document.body.offsetHeight - 300;
            tinymce.init({
                selector: '#tinymceEditer',
                branding: false,
                elementpath: false,
                height: height,
                language: 'zh_CN.GB2312',
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
                table_default_styles: {
                    width: '100%',
                    borderCollapse: 'collapse'
                },
                images_upload_url: "/ec-zx",
                setup: function (editor) {
                    editor.on('init', (e) => {
                        if (th.v)
                            tinymce.get('tinymceEditer').setContent(th.v);
                    });
                },
            });
        });
    }, destroyed() {
        tinymce.get('tinymceEditer').destroy();
    }
};
</script>
