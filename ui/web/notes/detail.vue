<template>
    <Row>
        <Col span="18">
            <card>
                <h2>{{notes.Title}}</h2>
                <small class="gray">
                    <router-link :to="'/p/member/'+notes.Username">{{notes.Username}}</router-link>
                    <!--· 1 天前 · 5194 次点击 &nbsp;-->
                </small>
            </card>
            <card>
                <div v-html="notes.Body"></div>
            </card>
        </Col>
        <Col span="6">
            <div style="background:#eee;padding:0 0 0 10px">
                <user-info></user-info>
            </div>
        </Col>
    </Row>
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
                    this.ajax('/gk-note/detail', {id: id}, (r, th) => {
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