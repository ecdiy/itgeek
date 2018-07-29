<template>
    <table cellpadding="0" cellspacing="0" border="0" width="100%">
        <tr>
            <td width="100" align="left">
                <x-button mini class="fl" v-if="no>1" @click.native="gop(no-1)">《上一页</x-button>
            </td>
            <td width="auto" align="center"><strong class="fade">{{no}}/{{all}}</strong></td>
            <td width="100" align="right">
                <x-button mini class="fr" v-if="no<all" @click.native="gop(no+1)">下一页》</x-button>
            </td>
        </tr>
    </table>
</template>

<script>
    export default {
        props: ['no', 'all', 'url'],
        methods: {
            gop(n) {
                vm.$router.push({path: this.url.replace('{}', n)});
            }
        }
    }
</script>
