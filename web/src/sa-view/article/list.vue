<template>
    <div class="vue-box">
        <div class="c-panel">
            <!-- 参数栏 -->
            <el-form :inline="true" size="mini">
                <el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-plus" @click="add">增加</el-button>
                </el-form-item>
                <el-form-item label="分类：" v-if="category_id == 7">
                    <el-cascader :options="categorys" :props="defaultProps" v-model="p.category_id"></el-cascader>
                </el-form-item>

                <el-form-item label="标题：">
                    <el-input type="text" v-model="p.title" placeholder="输入标题"></el-input>
                </el-form-item>

                <el-form-item style="min-width: 0px">
                    <el-button type="primary" icon="el-icon-search" @click="f5();">查询 </el-button>
                </el-form-item>
            </el-form>
            <!-- <div class="c-title">数据列表</div> -->
            <el-table :data="dataList" size="mini" ref="multipleTable">
                <el-table-column label="标题" prop="title"></el-table-column>
                <el-table-column label="分类" v-if="category_id == 7" prop="category_name"></el-table-column>
                <el-table-column label="图片">
                    <template slot-scope="s">
                        <el-image style="width:50px; height:50px"
                            :src="s.row.img ? s.row.img : require('@/static/img/no.png')"
                            :preview-src-list="[s.row.img ? s.row.img : require('@/static/img/no.png')]">
                        </el-image>
                    </template>
                </el-table-column>
                <el-table-column label="排序" prop="sort"></el-table-column>

                <el-table-column prop="address" label="操作" width="220px">
                    <template slot-scope="s">
                        <el-button class="c-btn" type="primary" icon="el-icon-edit" @click="update(s.row)">修改
                        </el-button>
                        <el-button class="c-btn" type="danger" icon="el-icon-delete" @click="del(s.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页 -->
            <div class="page-box">
                <el-pagination background layout="total, prev, pager, next, sizes, jumper" :current-page.sync="p.page"
                    :page-size.sync="p.pageSize" :total="dataCount" :page-sizes="[1, 10, 20, 30, 40, 50, 100]"
                    @current-change="f5(true)" @size-change="f5(true)">
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
            p: { pageSize: 10, page: 1, category_id: 0, title: "" },
            dataCount: 0,
            dataList: [],
            categorys: [],
            category_id: 0,
            defaultProps: {
                children: "childList",
                label: "name",
                value: "id",
                checkStrictly: true,
                emitPath: false,
            },
        };
    },
    methods: {
        // 数据刷新
        f5: function () {
            this.sa.put("/article/listPage", this.p).then((res) => {
                this.dataList = res.data.map((item) => {
                    return item;
                });
                this.dataCount = res.page_data.totalElements;
                //资质查询子类
                if (this.category_id == 7) {
                    let p = JSON.parse(JSON.stringify(this.p));
                    p.parent_id = 7;
                    this.sa.put("/category/listPage", p).then((res) => {
                        this.categorys = res.data
                        this.dataList = this.dataList.map(dto=>{
                            this.categorys.forEach(fl=>{
                                if(fl.id == dto.category_id){
                                    dto.category_name = fl.name;
                                }
                            })
                            return dto
                        })
                    })
                }
            });
        },
        // 删除
        del: function (data) {
            this.sa.confirm(
                "是否删除，此操作不可撤销",
                function () {
                    this.sa.delete("/article/" + data.id).then((res) => {
                        console.log(res);
                        this.sa.arrayDelete(this.dataList, data);
                        this.sa.ok(res.message);
                    });
                }.bind(this)
            );
        },



        //更新
        update(row) {
            row.type = this.category_id;
            row.categorys = this.categorys
            this.$refs["add-or-update"].open(row);
        },
        //添加
        add: function () {
            let row = { type: this.category_id }
            row.categorys = this.categorys
            this.$refs["add-or-update"].open(row);
        },
    },
    created: function () {
        // 7 资质，12 案例  16 主营业务
        this.category_id = parseInt(this.sa_admin.nativeTab.description);
        this.p.category_id = this.category_id
        this.f5();
    },
};
</script>

<style>
</style>
