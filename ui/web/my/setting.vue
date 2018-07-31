<template>

    <card>
        <Tabs size="small" :animated="false">
            <TabPane label="基本信息">
                <Form :labelWidth="80" style="width: 450px">
                    <FormItem label="编辑器">
                        <RadioGroup v-model="fm.EditType" type="button">
                            <Radio label="Markdown"></Radio>
                            <Radio label="HTML"></Radio>
                        </RadioGroup>
                    </FormItem>
                    <!--<FormItem label="笔记">-->
                    <!--<RadioGroup v-model="privacyMap.note" type="button">-->
                    <!--<Radio label="公开"></Radio>-->
                    <!--<Radio label="私有"></Radio>-->
                    <!--</RadioGroup>-->
                    <!--</FormItem>-->
                    <FormItem label="个人网站">
                        <Input v-model="info.HomeSite"/>
                    </FormItem>
                    <FormItem label="GitHub">
                        <Input v-model="info.GitHub"/>
                    </FormItem>
                    <FormItem label="签名">
                        <Input v-model="info.Sign"/>
                    </FormItem>
                    <FormItem label="所在公司">
                        <Input v-model="info.Company"/>
                    </FormItem>
                    <FormItem>
                        <Button @click="save">保存</Button>
                    </FormItem>
                </Form>
            </TabPane>
            <TabPane label="头像">
                <img :src="hdImg" border="0" style="max-height: 48px">
                <Upload :on-success="onSuccess" :show-upload-list="false"
                        action="/api/gk-upload/Avatar">
                    <Button icon="ios-cloud-upload-outline">上传头像</Button>
                </Upload>
            </TabPane>
            <TabPane label="修改密码">
                <Form :labelWidth="80" style="width: 450px">
                    <FormItem label="新密码">
                        <Input v-model="pass" type="password"/>
                    </FormItem>
                    <FormItem>
                        <Button @click="saveNewPass">保存</Button>
                    </FormItem>
                </Form>
            </TabPane>
        </Tabs>
    </card>

</template>

<script>
    export default {
        data() {
            return {
                pass: '',
                fm: {EditType: gk.user.EditType}, info: {}, privacy: 0, privacyMap: {note: '公开'}, v: 0
            }
        },
        computed: {

            hdImg() {
                return this.avatar(gk.user.Id) + this.v
            },
        }
        , methods: {
            saveNewPass() {
                this.ajax('/gk-user/setting/upPass', {Password: this.pass}, function (r, th) {
                    th.$Notice.success({
                        title: '设置密码成功'
                    });
                })
            },
            save() {
                var max = 255;
                if (this.privacyMap.note == '私有') {
                    this.privacy = (this.privacy | 1)
                } else {
                    this.privacy = (this.privacy & (max - 1))
                }
                window.gk.user.EditType = this.fm.EditType;
                let th = this;
                this.fm.Info = JSON.stringify(this.info)
                this.fm.Privacy = this.privacy
                this.ajax('/gk-user/setting/save', this.fm, function () {
                    th.$Notice.success({
                        title: '设置保存成功'
                    });
                })
            }, onSuccess() {
                window.avatarVer++;
                this.v++;
            },
            upload() {


            }
        }, created() {
            let th = this;
            this.ajax('/gk-user/setting/get', {}, function (r) {
                if (r.row.Info && r.row.Info.length > 0) {
                    th.info = JSON.parse(r.row.Info)
                }
                th.privacy = Number(r.row.Privacy);
                th.privacyMap.note = ((th.privacy & 1) == 1) ? '私有' : "公开"
            })
        }
    }
</script>
