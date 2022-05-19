<template>
    <el-dialog v-if="m" :title="title" :visible.sync="isShow" width="550px" top="10vh" :append-to-body="true"
        :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <!-- 参数栏 -->
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm"
                    label-width="120px">
                    <el-form-item label="分组" prop="group">
                        <el-input v-model="m.group"></el-input>
                    </el-form-item>
                    <el-form-item label="唯一标识" prop="code">
                        <el-input v-model="m.code"></el-input>
                    </el-form-item>
                    <el-form-item label="配置值" prop="value">
                        <el-input v-model="m.value"></el-input>
                    </el-form-item>
                    <el-form-item label="名称" prop="name">
                        <el-input v-model="m.name"></el-input>
                    </el-form-item>
                    <el-form-item label="类型">
                        <input-enum enumName="HtmlTypeEnum" v-model="m.type"></input-enum>
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
import inputEnum from "../../sa-resources/com-view/input-enum.vue";
export default {
    components: { inputEnum },
    props: ["params"],
    data() {
        return {
            m: {},
            title: "",
            isShow: false,
            rules: {
                group: [],
                code: [],
                value: [
                    {
                        min: 0,
                        max: 1000,
                        message: "长度在 1000 个字符",
                        trigger: "blur",
                    },
                ],
                name: [],
                type: [],
            },
            fileList: [],
        };
    },
    methods: {
        open: function (data) {
            this.isShow = true;
            if (data) {
                this.title = "修改 全局字典配置表";

                this.m = data;
            } else {
                this.m = { group: "", code: "", value: "", name: "", type: 0 };
                this.title = "添加 全局字典配置表";
            }
        },

        //提交全局字典配置表信息
        ok: function (formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.sa.post("/Config/save", this.m).then((res) => {
                        console.log(res);
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
    created() { },
};
</script>