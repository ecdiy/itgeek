<template>
    <div>
        <div>
            <x-input title="标题" v-model="fm.Title"></x-input>
        </div>
        <div>
            <group>
                <x-textarea title="内容" v-model="fm.Body"></x-textarea>
            </group>
        </div>
        <div style="margin: 0 15px;">
            <x-button @click.native="spp = true" type="primary">{{catTxt}}</x-button>
            <popup-picker :show.sync="spp" :data='catList' :show-cell="false" :fixed-columns="2" :columns=2
                          v-model='catId' @on-hide="changeCat"></popup-picker>
        </div>
        <div>
            <x-button @click.native="saveGoods">保存</x-button>
        </div>
    </div>
</template>

<script>
    import {XTextarea, PopupPicker} from 'vux'

    export default {
        components: {XTextarea, PopupPicker},
        data() {
            return {fm: {}, imgParam: {}, catList: [], catId: [], catTxt: '选择分类', spp: false};
        },
        methods: {
            cat(id) {
                for (var i = 0; i < this.catList.length; i++) {
                    if (this.catList[i].value == id) {
                        return this.catList[i]
                    }
                }
            },
            changeCat() {
                if (this.catId.length == 2) {
                    this.catTxt = '分类 : ' + this.cat(this.catId[0]).name + ' - ' + this.cat(this.catId[1]).name;
                }
            },
            init() {
                this.ajax('/gk-topic/luax/catgory/newNote', {}, (r,th) => {
                    th.catList = r.list;
                });
            }, saveGoods() {
                this.fm.CategoryId = this.catId[this.catId.length - 1];
                this.ajax("/gk-topic/luax/notes/save", this.fm, function (d,th) {
                    th.fm.Body = "";
                    th.fm.Title = "";
                    window.location = '/p/notes/list'
                });
            }, handleUploadImg(res, file, f) {
                this.url = file.response.uri;
                this.param.url = this.url;
                this.$emit('setImgData', this.paramName, this.param);
            }, selectHandle(selectedVal, selectedIndex, selectedText) {
                this.catText = '分类:' + selectedText.join(' => ');
                this.catId = selectedVal;
            },
        }, beforeRouteEnter(to, from, next) {
            next((v) => {
                v.init()
            })
        }
    }
</script>