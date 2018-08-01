<template>
    <gk-body>
        <div slot="left">

            <Breadcrumb>
                <BreadcrumbItem>
                    <go to="/">首页</go>
                </BreadcrumbItem>
                <BreadcrumbItem>
                    <go to="down">下载</go>
                </BreadcrumbItem>

            </Breadcrumb>
            <card>
                <div>
                    <div>{{detail.ResName}}</div>
                    <div>{{detail.ResDesc}}</div>
                </div>
            </card>
            <card>
                <Row>
                    <Col span="12">
                        <Rate disabled v-model="Rate"/>
                        综合评分：{{detail.Rate}}
                    </Col>
                    <Col span="12">
                        <div class="fr">
                            <span>
                            <Icon type="ios-star-outline" size="20"/>
                            收藏</span>
                            <span> <Icon type="ios-chatboxes-outline" size="20"/>
                            评论</span>
                            <span>  <Icon type="ios-alert-outline" size="20"/>
                                举报</span>
                        </div>

                    </Col>
                </Row>
            </card>
            <card>
                <Row>
                    <Col span="12">
                        所需: {{detail.ResScore}}铜币
                    </Col>
                    <Col span="12">
                        <Button size="large" class="fr" @click="down">立即下载</Button>
                    </Col>
                </Row>
            </card>
        </div>
        <div slot="right">
            <card>

            </card>
        </div>
    </gk-body>
</template>

<script>

    export default {
        data() {
            return {detail: {}, pn: []}
        }, computed: {
            Rate() {
                return Number(this.detail.Rate)
            }
        }, methods: {
            down() {
                window.location.href = "/api/gk-upload/resDown?Id=" + this.pn[1]
            }
        },
        created() {
            this.pn = window.location.pathname.split(",")
            this.ajax('/gk-upload/resDetail', {Id: this.pn[1], UserId: this.pn[2]})
        }
    }
</script>
