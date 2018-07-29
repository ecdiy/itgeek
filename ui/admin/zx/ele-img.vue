<template>
    <Row>
        <Col span="16">
            <Form :model="param" ref="form" :label-width="60">
                <FormItem label="图片地址">
                    <Input style="width: 250px" v-model="param.url"></Input>
                </FormItem>
                <FormItem label="图片链接">
                    <Input style="width: 250px" v-model="param.href"></Input>
                </FormItem>
                <FormItem label="图片宽" style="padding: 10px 0">
                    <Input style="width: 250px" v-model="param.width"></Input>
                </FormItem>
                <FormItem label="图片高">
                    <Input style="width: 250px" v-model="param.height"></Input>
                </FormItem>
            </Form>
        </Col>
        <Col span="8">
            <img style="max-height: 100px" :src="param.url" v-if="param.url!=''">
            <Upload action="/api/gk-admin/siteUpload" :show-upload-list="false" :data="param"
                    :on-success="handleUploadImg">
                <Button type="ghost" icon="ios-cloud-upload-outline">上传图片</Button>
            </Upload>
        </Col>
    </Row>
</template>

<script>
    export default {
        props: ['paramName', 'v'],
        created() {
            vm.$on("commitFormData", this.post)
        },
        data() {
            return {param: {url: ''}}
        }, watch: {
            paramName(vx) {
                if (!this.v) {
                    this.param = {url: ''}
                }
            }
            , v(vx) {
                if (vx) {
                    this.param = vx;
                } else {
                    this.param = {url: ''}
                }
            }
        }, methods: {
            post() {
                this.$emit('on-result-change', this.paramName, this.param);
            },
            handleUploadImg(r) {
                this.param.url = r.location;
                this.post()
            },
        }
    }
</script>
