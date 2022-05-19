<template>
    <div class="vue-box">
        <div class="c-panel">
            <!-- 参数栏 -->
            <el-form :inline="true" size="mini">

                <el-form-item label="请求ip：">
                    <el-input v-model="p.requestIp" placeholder="模糊查询"></el-input>
                </el-form-item>
                <el-form-item label="描述：">
                    <el-input v-model="p.description" placeholder="模糊查询"></el-input>
                </el-form-item>
                <el-form-item label="日志类型：">
                    <el-input v-model="p.logType" placeholder="模糊查询"></el-input>
                </el-form-item>
                <el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-search" @click="f5();">查询 </el-button>

                </el-form-item>
            </el-form>
            <!-- <div class="c-title">数据列表</div> -->
            <el-table :data="dataList" size="mini">
                <el-table-column label="方法名" prop="method" min-width="100" :show-overflow-tooltip="true">
                </el-table-column>
                <el-table-column label="用户" prop="username">
                </el-table-column>
                <el-table-column label="请求ip" prop="requestIp"></el-table-column>
                <el-table-column label="描述" prop="description"></el-table-column>
                <el-table-column label="创建时间" prop="createDateTime"></el-table-column>
                <el-table-column label="浏览器" prop="browser" min-width="100" :show-overflow-tooltip="true">
                </el-table-column>
                <el-table-column label="请求耗时" prop="time"></el-table-column>
                <el-table-column label="参数" prop="params" min-width="180" :show-overflow-tooltip="true">
                </el-table-column>
            </el-table>
            <!-- 分页 -->
            <div class="page-box">
                <el-pagination background layout="total, prev, pager, next, sizes, jumper" :current-page.sync="p.page" :page-size.sync="p.pageSize" :total="dataCount" :page-sizes="[1, 10, 20, 30, 40, 50, 100]" @current-change="f5(true)" @size-change="f5(true)">
                </el-pagination>
            </div>
        </div>

        <!-- 增改组件 -->
        <add-or-update ref="add-or-update"></add-or-update>
    </div>
</template>

<script>
//import inputEnum from "../../sa-resources/com-view/input-enum.vue";
import addOrUpdate from "./add.vue";
export default {
    components: {
        addOrUpdate,
        //inputEnum,
    },
    data() {
        return {
            p: {
                pageSize: 10,
                page: 1,
                requestIp: "",
                description: "",
                logType: "",
            },
            dataCount: 0,
            dataList: [],
        };
    },
    methods: {
        // 数据刷新
        f5: function () {
            this.sa.put("/Log/listPage", this.p).then((res) => {
                this.dataList = res.data.content.map((item) => {
                    return item;
                });
                this.dataCount = res.data.pageable.totalElements;
            });
        },
        // 删除
        del: function (data) {
            this.sa.confirm(
                "是否删除，此操作不可撤销",
                function () {
                    this.sa.delete("/Log/delete/" + data.id).then((res) => {
                        console.log(res);
                        this.sa.arrayDelete(this.dataList, data);
                        this.sa.ok(res.message);
                    });
                }.bind(this)
            );
        },

        exportFile: function () {
            this.p.pageSize = 1000;
            var titleObj = [
                { key: "requestIp", name: "请求ip" },
                { key: "address", name: "地址" },
                { key: "description", name: "描述" },
                { key: "browser", name: "浏览器" },
                { key: "time", name: "请求耗时" },
                { key: "method", name: "请求方式" },
                { key: "params", name: "参数" },
                { key: "logType", name: "日志类型" },
                { key: "exceptionDetail", name: "异常详情" },
                { key: "id", name: "id" },
            ];
            this.sa
                .exportFile("/Log/listPage", this.p, titleObj)
                .then((res) => {
                    console.log(res);
                });
        },

        //更新
        update(row) {
            this.$refs["add-or-update"].open(row);
        },
        //添加
        add: function () {
            this.$refs["add-or-update"].open(0);
        },
    },
    created: function () {
        this.f5();
    },
};
</script>

<style>
</style>
