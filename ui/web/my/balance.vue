<template>
    <card>
        <div>
            <Table :columns="column" :loading="loading" :data="list"></Table>
            <div style="clear: both;padding-top: 10px; ">
                <Page v-if="total>20" :total="total" :page-size="20" :current="current" show-total
                      @on-change="gop"></Page>
            </div>
        </div>
    </card>
</template>

<script>
    export default {
        data() {
            return {
                total: 0, list: [], loading: true, current: 1,
                column:
                    [{title: '日期', key: 'CreateAt', width: 150},
                        {title: '类型', key: 'ScoreType', width: 120},
                        {
                            title: '数额', width: 80, render: (h, p) => {
                                var st = '';
                                if (p.row.Fee < 0) {
                                    st = "color:red"
                                } else {
                                    st = "color:green"
                                }

                                return h('strong', {
                                    domProps: {style: st}
                                }, p.row.Fee)
                            }
                        },
                        {title: '余额', key: 'Score', width: 100},
                        {title: '描述', key: 'ScoreDesc', type: 'html'},]
            }
        }, methods: {
            gop(p) {
                vm.push({path: '/p/my/balance,' + p})
            }
        },
        created() {
            var pn = window.location.pathname;
            var p = pn.split(",")
            this.current = p[1]
            this.ajax('/gk-user/scoreLogList', {page: this.current})
        }
    }
</script>
