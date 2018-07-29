<template>
    <card>
        <base-page listUrl="/gk-admin/userList" :columns="column"></base-page>
    </card>
</template>

<script>
    import basePage from '../libs/base-table'
    export default {

        components: {basePage},
        data() {
            return {
                column: [{title: 'Username', key: 'Username'},
                    {title: '总活跃度', key: 'DauCount'},
                    {title: '今日活跃度', key: 'Dau'},
                    {title: '积分', key: 'Score'},]
            }
        }
    }
</script>
