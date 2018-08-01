<template>
    <gk-body>
        <div slot="left">
            <card>
                <ul>
                    <li v-for="it in list">
                        <go :to="'down/item,'+it.Id+','+it.UserId">{{it.ResName}}</go>
                    </li>
                </ul>
            </card>
        </div>
        <div slot="right">
            <card></card>
        </div>
    </gk-body>
</template>

<script>
    export default {
        data() {
            return {list: []}
        },
        created() {
            this.ajax('/gk-upload/resList', {page: 1})
        }

    }
</script>
