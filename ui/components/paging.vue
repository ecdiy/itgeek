<template>
    <div>
        <Table :columns="columns" :loading="loading" :data="list"></Table>
        <div style="clear: both;padding-top: 10px; ">
            <slot style="float: left"></slot>
            <Page v-if="total>ps" :total="total" style="text-align: right; padding-top: 10px; clear: both;"
                  :page-size="ps" show-sizer show-total
                  @on-change="load" @on-page-size-change="changPs"></Page>
        </div>
    </div>
</template>
<script>
    import Vue from 'vue'
    import Cookies from 'js-cookie';

    export default {
        props: ['columns', 'listUrl', 'name', 'autoLoad'],
        created() {
            if (this.name && this.name.length > 1) {
                var ps = Cookies.get(this.name + '_ps');
                if (ps && ps.length > 0) {
                    this.ps = Number(ps);
                }
            }
            if (!this.autoLoad || this.autoLoad == "")
                this.loadX(1, this.ps);
            VueOn(this.name + 'Item', (p) => {
                Vue.set(this.list[p.index], p.key, p.val);
            });
            VueOn(this.name + 'InsertItem', (p) => {
                this.list.push(p);
            });
            VueOn(this.name + 'DeleteItem', (p) => {
                this.list.splice(p, 1);
                this.total = this.total - 1;
            });
            VueOn(this.name + 'UpdateItem', (p) => {
                for (var k in p[1]) {
                    Vue.set(this.list[p[0]], k, p[1][k]);
                }
            });
            VueOn(this.name + 'PagingLoad', (p) => {
                this.loadX(1, this.ps, p);
            });
        }, data() {
            return {total: 0, ps: 10, p: 1, list: [], loading: true}
        }, methods: {
            load(data) {
                this.loadX(data, this.ps)
            }, loadX(page, pageSize, inp) {
                this.p = page;
                this.ps = pageSize;
                let th = this;
                var param = {Begin: (this.p - 1) * th.ps, PageSize: pageSize};
                if (inp) {
                    param = Object.assign(inp, param);
                }
                post(this.listUrl, param, function (data) {
                    th.loading = false;
                    if (data.Status && (!data.Status.Code || data.Status.Code == 0)) {
                        if (data.Result.r1) {
                            th.list = data.Result.r1;
                        } else {
                            th.list = [];
                        }
                        th.total = Number(data.Result.r0[0].Total);
                    }
                })
            }, changPs(nps) {
                this.ps = nps;
                if (this.name && this.name.length > 1) {
                    Cookies.set(this.name + '_ps', nps, {expires: 365, path: ''});
                }
                this.load(this.p, this.ps);
            }
        }
    }
</script>