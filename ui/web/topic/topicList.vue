<template>
    <div>
        <div v-for="it in list">
            <div class="row">
                <table cellpadding="0" cellspacing="0" border="0" width="100%">
                    <tr>
                        <td width="48" valign="top" align="center" v-if="self!=1">
                            <router-link :to="'/p/member/'+it.Username">
                                <img :src="avatar( it.UserId)" border="0"></router-link>
                        </td>
                        <td width="10"></td>
                        <td width="auto" valign="middle">
                        <span class="item_title"><router-link
                                :to="'/p/topic/detail,'+it.Id+','+it.UserId">{{it.Title}}</router-link></span>
                            <span class="topic_info"><div class="votes"></div>
                                    <router-link v-if="cId!=it.CategoryId" class="node"
                                                 @click="setId('0',it.CategoryId)"
                                                 :to="'/p/topic/list,0,'+it.CategoryId+',1'">{{it.CategoryName}}</router-link>
                                <strong v-if="self!=1">&nbsp;•&nbsp;
                                    <router-link :to="'/p/member/'+it.Username">{{it.Username}}</router-link>
                                </strong>
                                    &nbsp;•&nbsp; {{it.CreateAt}}
                                   <span v-if="it.ReplyCount>0">&nbsp;•&nbsp; 最后回复来自 <strong>
                                    <router-link :to="'/p/member/'+it.ReplyUsername">{{it.ReplyUsername}}</router-link>
                                    </strong></span>
                            </span>
                        </td>
                        <td width="70" align="right" valign="middle">
                            <router-link v-if="it.ReplyCount>0" :to="'/p/topic/detail,'+it.Id+','+it.UserId"
                                         class="count_livid">{{it.ReplyCount}}
                            </router-link>
                        </td>
                    </tr>
                </table>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        props: {list: Array, cId: "", self: ""}
    }
</script>
<style>
    .row {
        padding: 10px;
        font-size: 14px;
        line-height: 120%;
        text-align: left;
        border-bottom: 1px solid #e2e2e2;
    }

    .item_title {
        font-size: 16px;
        line-height: 130%;
        text-shadow: 0 1px 0 #fff;
        word-wrap: break-word;
        hyphens: auto;
        display: block;
        padding: 5px;
    }

    .topic_info {
        font-size: 12px;
        color: #ccc;
        line-height: 200%;
    }
</style>