<template>
    <div>
        <div class="cell"><h2>{{note.Title}}</h2></div>
        <div v-html="note.Body" class="cell"></div>
    </div>
</template>

<script>
    export default {
        data() {
            return {id: '', note: {}}
        }, methods: {
            init() {
                this.id = window.location.hash
                this.id = this.id.substr(7);
                this.ajax('/gk-topic/luax/notes/detail', {id: this.id}, (r,th) => {
                    th.note = r.Notes
                })
            }
        }, beforeRouteEnter(to, from, next) {
            next((v) => {
                v.init()
            })
        }
    }
</script>
