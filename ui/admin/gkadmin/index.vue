<template>
    <div>
        <Form :labelWidth="80" style="width: 400px">
            <FormItem label="网站名称">
                <Input v-model="fm.SiteName"/>
            </FormItem>
            <FormItem label="网站签名">
                <Input v-model="fm.Sign"/>
            </FormItem>
            <FormItem label="备案号">
                <Input v-model="fm.Miibeian"/>
            </FormItem>
            <FormItem label="LOGO">
                <div><img style="float: left" :src="fm.Logo" border="0" height="100">
                    <div style="float: left;padding-left: 60px">
                        <Upload style="float: left; " :on-success="onSuccess"
                                :show-upload-list="false"
                                :data="{'SiteId':SiteId}"
                                action="/api/gk-admin/siteUpload">
                            <Button type="ghost" icon="ios-cloud-upload-outline">上传Logo</Button>
                        </Upload>
                        <Icon style="float: left; clear: both;padding-top: 10px " @click="del" size="18" type="close"></Icon>
                    </div>
                </div>
            </FormItem>
            <FormItem>
                <Button @click="save">保存</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>
    export default {
        data() {
            return {fm: {SiteName: "", Logo: ""}, SiteId: window.gk.siteId}
        }, methods: {
            del() {
                this.fm.Logo = "";
            }, save() {
                this.ajax('/gk-admin/baseSave', {info: JSON.stringify(this.fm), key: "BaseInfo"}, (o, th) => {
                    th.$Notice.success({title: '设置保存成功'});
                })
            }, onSuccess(r) {
                this.fm.Logo = r.location;
            }
        }, created() {
            this.ajax("/gk-admin/baseGet", {key: "BaseInfo"}, (o, t) => {
                if (o.baseInfo && o.baseInfo.length > 2) {
                    t.fm = JSON.parse(o.baseInfo)
                }
            })
        }
    }
</script>
