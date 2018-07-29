<template>
    <card>
        <base-page listUrl="/gk-admin/topicList" :columns="column"></base-page>
    </card>
</template>

<script>

    export default {
        data() {
            return {column: [{title: '日期', key: 'CreateAt', width: 150},]}
        }
    }
</script>
