<template>
    <div style="background:#eee;padding: 10px 0;clear: both;position: relative">
        <Row>
            <Col span="18">
                <slot name="left"></slot>

                <card v-for="it in mainList">
                    <component :is="it.key" :param="it.val"></component>
                </card>
            </Col>
            <Col span="6">
                <div style="background:#eee;padding:0 0 0 10px">
                    <user-info v-if="gk.login"></user-info>
                    <login v-else></login>
                    <slot name="right"></slot>
                </div>

                <div v-for="it in rightList" style="background:#eee;padding:0 0 0 10px">
                    <card>
                        <component :is="it.key" :param="it.val"></component>
                    </card>
                </div>

            </Col>
        </Row>
    </div>
</template>

<script>
    export default {
        props: ['ecKey'],
        data() {
            return {gk: window.gk, rightList: [], mainList: []}
        }, created() {
            vm.$on("data", (p) => {
                this.gk = p;
            })
            if (this.ecKey) {
                this.ajax('/gk-topic/zxPage', {'PageKey': this.ecKey}, (r, th) => {
                    if (r.result && r.result.length > 2) {
                        var res = JSON.parse(r.result);
                        if (res.right && res.right.list) {
                            th.rightList = res.right.list;
                        }
                        if (res.main && res.main.list) {
                            th.mainList = res.main.list;
                        }
                    }

                })
            }
        }
    }
</script>
