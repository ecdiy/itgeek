<template>
    <Card>
        <table cellpadding="5" cellspacing="15" border="0" width="100%">
            <tr v-for="(it,idx) in list">
                <td width="60" valign="top" align="left">
                    <router-link :to="'/p/member/'+it.Username">
                        <img :src="avatar( it.UserId)" border="0"></router-link>
                </td>
                <td width="auto" align="left">
                    <h2 style="margin-bottom: 0px; margin-top: 0px; border-bottom: none;"><span
                            class="gray">{{idx+1}}.</span>
                        <router-link :to="'/p/member/'+it.Username">{{it.Username}}</router-link>
                    </h2>
                    <member-info :json="it.Info"></member-info>
                </td>
                <td width="200" align="center">
                    <div class="balance_area" style="font-size: 24px; line-height: 24px;">&nbsp;{{it.Dau}}&nbsp;</div>
                </td>
            </tr>
        </table>
    </Card>
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
