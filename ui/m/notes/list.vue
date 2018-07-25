<template>
    <div class="bar">
        <flexbox>
            <flexbox-item>
                <x-button link="/notes/new">新建笔记</x-button>
            </flexbox-item>
            <flexbox-item>
                <x-button link="/notes/category">分类管理</x-button>
            </flexbox-item>
        </flexbox>
        <group v-for="c in catList" v-if="c.ParentId=='0'">
            <cell :title="c.Name">
                <x-icon slot="icon" type="android-folder-open"></x-icon>
            </cell>
            <cell @click.native="setOpen(c2.Id)" class="sub" :title="c2.Name" v-for="c2 in catList"
                  :style="catOpen==c2.Id? 'color:green;':''"
                  v-if="c2.ParentId==c.Id && c2.ItemCount>0">
                <x-icon slot="icon" type="folder"></x-icon>
                {{c2.ItemCount}}
            </cell>
        </group>
        <group>
            <cell @click.native="setOpen('0')" title="未分类"
                  :style="catOpen=='0'? 'color:green;':''">
                <x-icon slot="icon" type="folder"></x-icon>
            </cell>
            <cell v-for="it in notList" v-if="it.CategoryId==catOpen" :link="'/note/'+it.Id" :title="it.Title">
                <x-icon slot="icon" type="document"></x-icon>
            </cell>
        </group>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                notList: [], catList: [], subShow: false, catOpen: "0"
            };
        }, methods: {
            setOpen(i) {
                this.catOpen = i;
            }
        }, beforeRouteEnter(to, from, next) {
            next((v) => {
                this.ajax('/gk-topic/luax/notes/list', {}, (r) => {
                    v.notList = r.list;
                    v.catList = r.cat;
                });
            })
        }
    }
</script>
<style scoped>
    .sub {
        padding-left: 40px;
    }
</style>