<template>
    <div class="box">
        <div>
            <table cellpadding="0" cellspacing="0" border="0" width="100%">
                <tr>
                    <th>日期/类型</th>
                    <th>数额/余额</th>
                    <th>描述</th>
                </tr>
                <tr v-for="it in list">
                    <td>
                        {{it.CreateAt}}<br>{{it.ScoreType}}
                    </td>
                    <td>
                        {{it.Fee}}<br>{{it.Score}}
                    </td>
                    <td v-html="it.ScoreDesc">
                    </td>
                </tr>
            </table>
        </div>
        <page :all="totalPage" :no="current" :url="'/p/my/balance,{}'"></page>
    </div>
</template>

<script>
    export default {
        data() {
            return {list: [], total: 0, current: 1}
        }, computed: {
            totalPage() {
                return Math.ceil(this.total / 20)
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

<style scoped>
    td, th {
        border-bottom: 1px solid #e9eaec;
        padding: 3px;
    }
    th{padding: 5px}
</style>