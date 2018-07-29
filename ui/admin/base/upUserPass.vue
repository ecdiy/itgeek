<template>
    <card>
        <div style="width:300px;margin: auto">
            <div class="login-con">
                <Card :bordered="false">

                    <div class="form-con">
                        <Form :model="form" ref="form">
                            <FormItem prop="SiteAdminUser" :rules="rules.SiteAdminUser" label="用户名"
                                      :error="err.SiteAdminUser">
                                <Input v-model="form.SiteAdminUser" placeholder="请输入用户名">
                                <span slot="prepend">
                                    <Icon :size="16" type="person"></Icon>
                                </span>
                                </Input>
                            </FormItem>
                            <FormItem prop="SiteAdminPass" :rules="rules.SiteAdminPass" label="密码"
                                      :error="err.SiteAdminPass">
                                <Input type="password" v-model="form.SiteAdminPass" placeholder="请输入密码">
                                <span slot="prepend">
                                    <Icon :size="14" type="locked"></Icon>
                                </span>
                                </Input>
                            </FormItem>

                            <FormItem>
                                <Button @click="handleSubmit" type="primary" long>修改</Button>
                            </FormItem>
                        </Form>
                    </div>
                </Card>
            </div>
        </div>
    </card>
</template>


<script>

    export default {
        data() {
            return {

                form: {SiteAdminUser: ''},
                err: {SiteAdminUser: "", SiteAdminPass: ""},
                rules: {
                    SiteAdminUser: [
                        {required: true, min: 3, message: '账号不能为空，长度至少5位', trigger: 'blur'}
                    ],
                    SiteAdminPass: [
                        {required: true, min: 6, message: '密码不能为空，长度至少6位', trigger: 'blur'}
                    ]
                }
            };
        },

        methods: {
            handleSubmit() {
                for (var k in this.err) {
                    this.err[k] = "";
                }
                this.$refs.form.validate((valid) => {
                    if (valid) {
                        this.ajax('/gk-admin/newAdminUser', this.form, (r, th) => {
                            th.$Notice.success({title: '设置保存成功'});
                        });
                    }
                });
            }
        }
    };
</script>