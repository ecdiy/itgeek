<template>
    <gk-body>
        <div slot="left">
            <card>
                <Row>
                    <Col span="18">

                        <Form :labelWidth="80" style="width: 450px">
                            <FormItem label="资源类型">
                                <Upload :on-success="onSuccess" :show-upload-list="false"
                                        action="/api/gk-upload/Avatar">
                                    <Button icon="ios-cloud-upload-outline">选择文件</Button>
                                </Upload>
                            </FormItem>

                            <FormItem label="资源类型">
                                <Select v-model="fm.ResType" style="width:200px">
                                    <Option value="1">文档</Option>
                                    <Option value="2">代码</Option>
                                    <Option value="3">工具</Option>
                                    <Option value="4">其它</Option>
                                </Select>
                            </FormItem>
                            <FormItem label="所属分类">
                                <Select v-model="fm.CategoryId" style="width:200px">
                                    <Option value="">请选择</Option>
                                    <Option value="21">移动开发</Option>
                                    <Option value="16">开发技术</Option>
                                    <Option value="15">课程资源</Option>
                                    <Option value="19">网络技术</Option>
                                    <Option value="12">操作系统</Option>
                                    <Option value="10">安全技术</Option>
                                    <Option value="18">数据库</Option>
                                    <Option value="14">行业</Option>
                                    <Option value="13">服务器应用</Option>
                                    <Option value="20">信息化</Option>
                                    <Option value="11">存储</Option>
                                    <Option value="17">考试认证</Option>
                                    <Option value="22">云计算</Option>
                                    <Option value="23">大数据</Option>
                                    <Option value="24">跨平台</Option>
                                    <Option value="25">音视频</Option>
                                    <Option value="26">游戏开发</Option>
                                    <Option value="27">人工智能</Option>
                                    <Option value="28">区块链</Option>
                                    <Option value="29">其它</Option>
                                </Select>
                            </FormItem>
                            <FormItem label="资源分数">
                                <Input/>
                            </FormItem>
                            <FormItem label="资源标签">
                                <Input/>
                            </FormItem>
                            <FormItem label="资源描述">
                                <Input type="textarea" :autosize="{minRows: 2,maxRows: 5}"/>
                            </FormItem>
                            <Button icon="ios-cloud-upload-outline">发布</Button>
                        </Form>
                    </Col>
                    <Col span="6">
                        <card>
                <span slot="title">
                    上传须知
                </span>
                            <ul style="padding:5px 10px">
                                <li>扣除平台服务积分，1M 1积分，不足1M按1M计算</li>
                                <li>如涉及侵权内容,您的资源将被移除</li>
                                <li>请勿上传小说、mp3、图片等与技术无关的内容.一旦发现将被删除</li>
                                <li>请勿在未经授权的情况下上传任何涉及著作权侵权的资源，除非该资源完全由您个人创作</li>
                                <li>点击上传资源即表示您确认该资源不违反资源分享的使用条款，并且您拥有该资源的所有版权或者上传资源的授权</li>
                                <li>您上传的资源如果因版权、使用、内容完整度 等原因被举报并通过官方审核，将扣除通过该资源获得的全部积分</li>
                            </ul>
                        </card>
                    </Col>
                </Row>
            </card>
        </div>
        <div slot="right">


        </div>
    </gk-body>
</template>

<script>
    export default {
        data() {
            return {fm: {}}
        }, methods: {
            onSuccess() {
                window.avatarVer++;
                this.v++;
            },
        }
    }
</script>
