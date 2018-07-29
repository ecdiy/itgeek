<template>
    <div>
        <div class="cell swipe-wrapper" v-for="(it,index) in list" :key="it.Id">
            <table cellpadding="0" cellspacing="0" border="0" width="100%">
                <tbody>
                <tr>
                    <td width="24" valign="top" align="center">
                        <router-link :to="'/p/member/'+it.Username"><img
                                :src="avatar(it.UserId)"
                                border="0" align="default"
                                style="max-width: 24px; max-height: 24px;"></router-link>
                    </td>
                    <td width="10"></td>
                    <td width="auto" valign="middle"><span class="small fade">
                            <go :to="'/p/topic/list,0,'+it.CategoryId+',1'" class="node">{{it.CategoryName}}</go> &nbsp;•&nbsp; <strong><router-link
                            :to="'/p/member/'+it.Username">{{it.Username}}</router-link></strong></span>
                        <div>
                                <span class="item_title"><router-link
                                        :to="'/p/topic/detail,'+it.Id+','+it.UserId">{{it.Title}}</router-link></span>
                        </div>
                        <div>
                            <span class="small fade">{{ it.CreateAt}}</span>
                            <span class="small fade" v-if="it.ReplyUsername!=''">
                             &nbsp;•&nbsp;最后回复 <strong><router-link
                                    :to="'/p/member/'+it.ReplyUsername">{{it.ReplyUsername}}</router-link></strong></span>
                        </div>
                    </td>
                    <td width="70" align="right" valign="middle"> {{it.ReplyCount}}</td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
    export default {
        props: {list: Array, cId: "", self: ""}
    }
</script>
