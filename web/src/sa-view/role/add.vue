<template>
    <el-dialog v-if="m" :title="title" :visible.sync="isShow" width="550px" top="10vh" :append-to-body="true" :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <!-- 参数栏 -->
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm" label-width="120px">

                    <el-form-item label="角色名称" prop="name">
                            <el-input v-model="m.name"></el-input>
                    </el-form-item>

                      <el-form-item label="角色编码" prop="code">
                            <el-input v-model="m.code" :disabled="m.code? true: false"  ></el-input>
                    </el-form-item>

                    <el-form-item label="描述" prop="description">
                        <el-input type="textarea" rows="2" placeholder="描述" v-model="m.description"></el-input>
                    </el-form-item>
                    <el-form-item label="排序" prop="sort">
                        <el-input v-model="m.sort" type="number"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <span class="c-label">&emsp;</span>
                        <el-button type="primary" icon="el-icon-plus" size="small" @click="ok('ruleForm')">确定
                        </el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </el-dialog>
</template>

<script>
import selectData from "../../sa-resources/com-view/select-data.vue";
export default {
    components: { selectData },
    //components: { inputEnum },
    props: ["params"],
    data() {
        return {
            m: {sort: 0},
            title: "",
            isShow: false,
            rules: {
                name: [
                    {
                        required: true,
                        message: "请输入部门名称",
                        trigger: "blur",
                    },
                ],
                description: [],
            },
            fileList: [],
        };
    },
    methods: {
        open: function (data) {
            this.isShow = true;
            if (data) {
                this.title = "修改 部门";

                this.m = data;
            } else {
                this.m = { name: "", description: "" , sort:1};
                this.title = "添加 部门";
            }
        },

        //提交部门信息
        ok: function (formName) {
            console.log(formName)
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.m.sort = parseInt(this.m.sort)
                    this.sa.post("/role/save", this.m).then((res) => {
                        this.$parent.f5();
                        this.isShow = false;
                    });
                } else {
                    console.log("error submit!!");
                    return false;
                }
            });
        },
    },
    created() {},
};
</script>