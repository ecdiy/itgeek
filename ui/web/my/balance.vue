<template>
    <card>
        <base-page listUrl="/gk-user/scoreLogList" :columns="column"></base-page>
    </card>
</template>

<script>
    export default {
        data() {
            return {
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
        }
    }
</script>
