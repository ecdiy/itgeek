<template>
    <gk-body>
        <div slot="left">
            <card>
                上传权限： 50M
                <br>
                <table width="400" cellpadding="5" cellspacing="15" border="0">
                    <tr>
                        <td>上传资源</td>
                        <td>下载资源</td>
                    </tr>
                    <tr>
                        <td>9</td>
                        <td>9</td>
                    </tr>
                </table>

            </card>

            <card>
                <Tabs size="small" :animated="false">
                    <TabPane label="上传资源">
                        <go to="down/upload">
                            <Button type="success" size="large">上传资源</Button>
                        </go>
                    </TabPane>

                    <TabPane label="下载明细">
                    </TabPane>
                    <TabPane label="收藏">
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
    export default {}
</script>

<style scoped>

</style>