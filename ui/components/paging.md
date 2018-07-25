``` 
<div id="app">
    <app-kushow></app-kushow>
</div>
<script type="text/x-template" id="app-kushow">
    <div> 
        <paging list-url="/sysuser/UserSp/OrderList" :columns="columns" name="EcOrderList"></paging>
    </div>
</script>


import paging from '../../components/paging';

Vue.component('app-kushow', {
    template: '#app-kushow',
    components: {paging}, data() {
        return {
            columns: [{title: '订单Id', key: 'OrderId'}, {title: 'UserId', key: 'UserId'},
                {title: 'Username', key: 'Username'}, {title: '有效期', key: 'ValidDay'},
            ]
        }
    }
});

```
通讯接口，改变list,index项值：
 Hub.$emit('item', {index: this.row.index, key: "WareStatus", val: 1});