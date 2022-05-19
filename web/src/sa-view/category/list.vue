<template>
    <div class="vue-box">
        <div class="c-panel">
            <!-- 参数栏 -->
            <el-form :inline="true" size="mini">
                <el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-plus" @click="add(0)">增加</el-button>
                    <el-button type="primary" icon="el-icon-download" @click="exportFile()">导出</el-button>
                </el-form-item>
            </el-form>
            <!-- <div class="c-title">数据列表</div> -->
            <el-table :data="dataList" size="mini" row-key="id" default-expand-all :tree-props="{children: 'childList', hasChildren: 'hasChildren'}">
                
                <el-table-column label="分类名称" prop="name"></el-table-column>
                <el-table-column label="编号" prop="id"></el-table-column>
                <el-table-column v-if="p.type!=3" label="链接" prop="url"></el-table-column>
                <el-table-column label="排序" prop="sort"></el-table-column>
                <el-table-column prop="address" label="操作" width="280px">
                    <template slot-scope="s">
                        <el-button class="c-btn" type="primary" icon="el-icon-edit" @click="update(s.row)">修改</el-button>
                           <el-button class="c-btn" type="warning" v-if="s.row.parent_id == 0 && s.row.id != 1" icon="el-icon-delete" @click="add(s.row)">添加子分类</el-button>
                        <el-button class="c-btn" type="danger" v-if="s.row.parent_id != 0" icon="el-icon-delete" @click="del(s.row)">删除</el-button>
                    </template>
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
import sa_admin_code_util from "./../../sa-resources/index/admin-util.js"; // admin代码util
export default {
    components: {
        addOrUpdate,
    },
    data() {
        return {
            p: {
                pageSize: 50,
                page: 1,
                type: 0,
            },
            dataCount: 0,
            dataList: [],
        };
    },
    methods: {
        // 数据刷新
        f5: function () {
            this.sa.put("/category/listPage", this.p).then((res) => {
                this.dataList = sa_admin_code_util.arrayToTree(
                    res.data
                ); // 一维转tree
                console.log(8888, this.dataList)
                this.dataCount = res.page_data.totalElements;
            });
        },
        // 删除
        del: function (data) {
            this.sa.confirm(
                "是否删除，此操作不可撤销",
                function () {
                    this.sa.delete("/category/" + data.id).then((res) => {
                        console.log(res);
                        this.sa.arrayDelete(this.dataList, data);
                        this.sa.ok(res.message);
                        this.f5()
                    });
                }.bind(this)
            );
        },


        //更新
        update(row) {
            row.type = this.p.type
            this.$refs["add-or-update"].open(row);
        },
        //添加
        add: function (row) {
            if(row){
                row = {parent_id: row.id, parentName:row.name, type: this.p.type};
            }else{
                 row = {parent_id: 0, parentName:'顶级', type: this.p.type};
            }
            this.$refs["add-or-update"].open(row);
        },
        // 修改菜单权限
        menu_setup: function (data) {
            this.$refs["menu-setup"].show(data.id, data.name);
        },
    },
    created: function () {
        this.p.type = parseInt(this.sa_admin.nativeTab.description);
        this.f5();
    },
};
</script>

<style>
</style>
