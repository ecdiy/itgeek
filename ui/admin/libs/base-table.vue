<template>
    <div>
        <Table :columns="columns" :loading="loading" :data="list"></Table>
        <div style="clear: both;padding-top: 10px; ">
            <Page v-if="total>20" :total="total" :page-size="20" :current="current" show-total @on-change="load"></Page>
        </div>
    </div>
</template>
<script>
    export default {
        props: ['columns', 'listUrl'],
        created() {
            this.load(1);
        }, data() {
            return {total: 0, list: [], loading: true, current: 1}
        }, methods: {
            load(p) {
                this.loading = true;
                this.current = p;
                this.ajax(this.listUrl, {page: this.current})
            }
        }
    }
</script>