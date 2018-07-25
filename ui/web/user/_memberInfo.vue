<template>
    <div>
        <a :href="info.HomeSite" target="_blank" v-if="info.HomeSite && info.HomeSite.length>0">
            <Icon size="20" type="ios-home-outline"></Icon>
            {{info.HomeSite}}</a>

        <a :href="info.GitHub" target="_blank" v-if="info.GitHub && info.GitHub.length>0">
            <Icon size="20" type="social-github"></Icon>
            {{info.GitHub}}</a>
    </div>
</template>

<script>
    export default {
        props: ["json"],
        computed: {
            info() {
                return this.json && this.json.length > 2 ? JSON.parse(this.json) : {};
            }
        }
    }
</script>
