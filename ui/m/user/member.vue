<template>
    <div>

        <h2>{{username}}</h2>

    </div>
</template>

<script>
    export default {
        data() {
            return {
                user: {}, username: "",
            };
        },
        methods: {
            init() {
                var id = window.location.hash
                if (id.length > 9) {
                    this.username = id.substr(9);

                }
            }
        },
        mounted() {
            this.init();
        },
        activated() {
            this.init();
        }
    }
</script>

<style scoped>

</style>