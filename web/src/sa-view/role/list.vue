<template>
    <div class="vue-box">
        <div class="c-panel">
            <!-- 参数栏 -->
            <el-form :inline="true" size="mini">
                <el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-plus" @click="add">增加</el-button>
                    <el-button type="primary" icon="el-icon-download" @click="exportFile()">导出</el-button>
                </el-form-item>
                <el-form-item label="部门名称：">
                    <el-input v-model="p.name" placeholder="模糊查询"></el-input>
                </el-form-item>
                <el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-search" @click="f5();">查询 </el-button>

                </el-form-item>
            </el-form>
            <!-- <div class="c-title">数据列表</div> -->
            <el-table :data="dataList" size="mini" row-key="id"  :tree-props="{children: 'childList', hasChildren: 'hasChildren'}">
                <!--start-- 需要时候可以随时开启
				<el-table-column type="selection"></el-table-column>
				<el-table-column label="Id" prop="id"></el-table-column>
				-->
                <el-table-column label="部门名称" prop="name"></el-table-column>
				<el-table-column label="部门编码" prop="code"></el-table-column>
                <el-table-column label="描述" prop="description"></el-table-column>
                <el-table-column label="排序" prop="sort"></el-table-column>

                <el-table-column prop="address" label="操作" width="280px">
                    <template slot-scope="s">
                        <el-button class="c-btn" type="success" icon="el-icon-finished" @click="menu_setup(s.row)">权限分配</el-button>
                        <el-button class="c-btn" type="primary" icon="el-icon-edit" @click="update(s.row)">修改</el-button>
                        <el-button class="c-btn" type="danger" v-if="s.row.creatorId != 0" icon="el-icon-delete" @click="del(s.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页 -->
            <div class="page-box">
                <el-pagination background layout="total, prev, pager, next, sizes, jumper" :current-page.sync="p.page" :page-size.sync="p.pageSize" :total="dataCount" :page-sizes="[1, 10, 20, 30, 40, 50, 100]" @current-change="f5(true)" @size-change="f5(true)">
                </el-pagination>
            </div>
        </div>
        <!-- 设置权限  -->
        <menu-setup ref="menu-setup"></menu-setup>
        <!-- 增改组件 -->
        <add-or-update ref="add-or-update"></add-or-update>
    </div>
</template>

<script>
//import inputEnum from "../../sa-resources/com-view/input-enum.vue";
import addOrUpdate from "./add.vue";
import menuSetup from "./menu-setup.vue"; // 设置权限
import sa_admin_code_util from "./../../sa-resources/index/admin-util.js"; // admin代码util
export default {
    components: {
        addOrUpdate,
        menuSetup,
        //inputEnum,
    },
    data() {
        return {
            p: {
                pageSize: 10,
                page: 1,
                name: "",
            },
            dataCount: 0,
            dataList: [],
        };
    },
    methods: {
        // 数据刷新
        f5: function () {
            this.sa.put("/role/listPage", this.p).then((res) => {
                this.dataList = sa_admin_code_util.arrayToTree(
                    res.data
                ); // 一维转tree
                this.contentList = res.data;
                this.dataCount = res.page_data.totalElements;
            });
        },
        // 删除
        del: function (data) {
            this.sa.confirm(
                "是否删除，此操作不可撤销",
                function () {
                    this.sa.delete("/role/" + data.id).then((res) => {
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
                {
                    key: "name",
                    name: "部门名称",
                },
                {
                    key: "description",
                    name: "描述",
                },
                {
                    key: "sort",
                    name: "排序",
                },
                {
                    key: "id",
                    name: "id",
                },
            ];
            this.sa
                .exportFile("/Attachment/listPage", this.p, titleObj)
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
        // 修改菜单权限
        menu_setup: function (data) {
            this.$refs["menu-setup"].show(data.id, data.name);
        },
    },
    created: function () {
        this.f5();
    },
};
</script>

<style>
</style>
