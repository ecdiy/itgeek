<template>
    <gk-body>
        <div slot="left">
            <card>
                上传权限：{{info.UpLimit}}M
                <br>
                <table width="400" cellpadding="5" cellspacing="15" border="0">
                    <tr>
                        <td>上传资源</td>
                        <td>下载资源</td>
                    </tr>
                    <tr>
                        <td>{{info.UploadItem}}</td>
                        <td>{{info.DownItem}}</td>
                    </tr>
                </table>
            </card>

            <card>
                <Tabs size="small" :animated="false">
                    <TabPane label="上传资源">

                        <ul>
                            <li v-for="it in list">
                                <go :to="'down/item,'+it.Id+','+ UserId">{{it.ResName}}</go>
                                <div v-html="it.ResDesc"></div>
                                <div>
                                    上传时间:{{it.CreateAt}}
                                </div>
                            </li>
                        </ul>

                        <go to="down/upload">
                            <Button type="success" size="large">上传资源</Button>
                        </go>
                    </TabPane>

                    <TabPane label="下载明细">
                    </TabPane>
                    <TabPane label="收藏">
                        <ul>
                            <li></li>
                        </ul>
                    </TabPane>
                </Tabs>
            </card>
        </div>
        <div slot="right">
            <card></card>
        </div>
    </gk-body>
</template>

<script>
    export default {
        data() {
            return {
                info: {}, list: [], fav: [], UserId: gk.user.Id,
            }
        },
        created() {
            this.ajax('/gk-upload/resInfo', {})
        }
    }
</script>

<style scoped>

</style>