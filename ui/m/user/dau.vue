<template>
    <div class="box">
        <table cellpadding="5" cellspacing="15" border="0" width="100%">
            <tr v-for="(it,idx) in list">
                <td width="50" valign="top" align="left">
                    <router-link :to="'/p/member/'+it.Username">
                        <img :src="avatar( it.UserId)" border="0"></router-link>
                </td>
                <td width="auto" align="left">
                     <span
                             class="gray">{{idx+1}}.</span>
                    <router-link :to="'/p/member/'+it.Username">{{it.Username}}</router-link>

                    <member-info :json="it.Info"></member-info>
                </td>
                <td width="50" align="center">
                    {{it.Dau}}
                </td>
            </tr>
        </table>
    </div>
</template>

<script>
    import memberInfo from './_memberInfo'

    export default {
        components: {memberInfo},
        data() {
            return {list: []}
        },
        created() {
            this.ajax('/gk-user/topDau')
        }
    }
</script>
