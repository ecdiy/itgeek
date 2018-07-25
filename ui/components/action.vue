<template>
  <Modal :width="modalWidth" v-model="fmModal" :title="actionTitle+name " :mask-closable="false" :closable="true">
    <slot style="float: left"></slot>
    <div slot="footer">
      <span style="text-align: left;float: left">{{countInfo}}</span>
      <Button type="ghost" size="large" @click="undo">取消</Button>
      <Button v-if="action==='update' && canDel" type="ghost" size="large" @click="del">删除{{name}}</Button>
      <Button type="primary" size="large" @click="postForm">{{actionTitle}}</Button>
    </div>
  </Modal>
</template>

<script>

  export default {
    props: ['name', 'tbl', 'pk', 'actionUrl', 'modalWidth'],
    data() {
      return {
        countInfo: "", action: "", actionTitle: "", fmModal: false, canDel: false, obj: {}, fm: {},
      }
    }, created() {
      VueOn(this.tbl + 'OpenInsert', () => {
        this.action = "insert";
        this.fmModal = true;
        this.actionTitle = "添加";
      });
      VueOn(this.tbl + 'OpenUpdate', (p) => {
        this.action = "update";
        this.fmModal = true;
        this.actionTitle = "修改";
        this.obj = p[0];
        this.canDel = p[1];
      });

      VueOn(this.tbl + 'CountInfo', (p) => {
        this.countInfo = p;
      });

      VueOn(this.tbl + 'FmData', this.formAction);
    }, methods: {
      formAction(fm) {
        fm.table = this.tbl;
        fm.pk = this.pk;
        fm.action = this.action;
        let th = this;
        this.fm = fm;
        for (var k in fm) {
          if (k.indexOf("Date") > 0 && Object.prototype.toString.call(fm[k]) === "[object Date]") {
            fm[k] = fm[k].Format("yyyy-MM-dd");
          }
        }
        post(th.actionUrl, fm, function (data) {
          th.fmModal = false;
          if (data.Status && (!data.Status.Code || data.Status.Code === 0)) {
            if (th.action === "insert") {
              var obj = {};
              Object.assign(obj, th.fm);
              obj[th.pk] = data.Result;
              Hub.$emit(th.tbl + 'InsertItem', obj);
            }
            if (th.action === "update") {
              Hub.$emit(th.tbl + 'UpdateItem', [th.obj.index, th.fm]);
            }
            if (th.action === "delete") {
              Hub.$emit(th.tbl + 'DeleteItem', th.obj.index);
            }
          } else {
            if (data.Status && data.Status.Code === 12) {
              th.$Notice.error({title: "数据约束，操作失败！", desc: data.Result});
            }
          }
        })
      }, del() {
        this.$Modal.confirm({
          title: '删除', content: '<p></p>',
          onOk: () => {
            this.action = "delete";
            var r = {};
            r[this.pk] = this.obj.row[this.pk];
            this.formAction(r);
          }
        });
      }, undo() {
        this.fmModal = false;
      }, postForm() {
        Hub.$emit(this.tbl + 'GetFmData', {});
      }
    }
  }
</script>
