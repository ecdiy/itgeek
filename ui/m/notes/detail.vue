<template>
    <div>

        <div>
            <h2>{{notes.Title}}</h2>
            <small class="gray"><router-link :to="'/p/member/'+notes.Username">{{notes.Username}}</router-link>
                <!--· 1 天前 · 5194 次点击 &nbsp;-->
            </small>
        </div>
        <div class="cell" v-html="notes.Body"></div>
    </div>
</template>

<script>
    export default {
        data() {
            return {notes: {},};
        }, methods: {
            init() {
                var id = window.location.hash
                if (id.length > 13) {
                    id = id.substr(13);
                    this.ajax('/gk-topic/luax/notes/detail', {id: id}, (r,th) => {
                        th.notes = r.Notes
                    })
                }
            }
        }, mounted() {
            this.init();
        }, activated() {
            this.init();
        }
    }
</script>