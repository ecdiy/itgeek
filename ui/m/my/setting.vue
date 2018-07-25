<template>

    <Tabs size="small" :animated="false">
        <TabPane label="基本信息">

                <group>
                    <FormItem label="编辑器">
                        <RadioGroup v-model="fm.EditType" type="button">
                            <Radio label="Markdown" :true-value="0"></Radio>
                            <Radio label="HTML" :true-value="1"></Radio>
                        </RadioGroup>
                    </FormItem>
                </group>
                <FormItem>
                    <Button @click="save">保存</Button>
                </FormItem>

        </TabPane>

        <TabPane label="头像">
            <img :src="hdImg" border="0">
            <Upload :on-success="onSuccess" :show-upload-list="false"
                    action="/gk-upload/Avatar">
                <Button type="ghost" icon="ios-cloud-upload-outline">上传头像</Button>
            </Upload>
        </TabPane>
    </Tabs>


</template>

<script>
    export default {
        data() {
            return {
                fm: {EditType: gk.user.EditType},
                hdImg: '/avatar/' + gk.user.Id + '.png?' + (new Date()).getTime()
            }
        }, methods: {
            save() {
                window.gk.user.EditType = this.fm.EditType;
                this.ajax('/gk-user/setting/save', this.fm, function (r,th) {
                    th.$Notice.success({
                        title: '设置保存成功'
                    });
                })
            }, onSuccess() {
                window.avatarVer++;
                this.hdImg = '/avatar/' + gk.user.Id + '.png?' + (new Date()).getTime()
            },
            upload() {


            }
        }
    }
</script>
