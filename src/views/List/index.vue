<template>
    <div class="make-order-list">
        <h2>预约列表</h2>
        <el-table :data="tableData" style="width: 100%;height:100%">
        <el-table-column prop="CreatedAt" label="日期" width="250"></el-table-column>
        <el-table-column prop="phoneNum" label="手机号" width="180"></el-table-column>
        <el-table-column prop="city_name" label="地区"></el-table-column>
    </el-table>
    </div>
</template>

<script>
import moment from "moment"
import { getOrders } from "../../api/list"
export default {
    data() {
        return {
            tableData: []
        }
    },
    created(){
        getOrders().then(res => {
            const syncData = res.data.map(item => {
                return {
                    ...item,
                    CreatedAt:moment(item.CreatedAt).format("YYYY-MM-DD hh:mm:ss")
                }
            })
            this.tableData = syncData
        })
    }
}
</script>
<style>
.make-order-list{
    padding: 10px;
}
</style>