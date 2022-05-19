<template>
    <el-dialog v-if="m" :title="title" :visible.sync="isShow" width="550px" top="10vh" :append-to-body="true" :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <!-- 参数栏 -->
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm" label-width="120px">
                    <el-form-item label="请求ip" prop="requestIp">
                        <el-input v-model="m.requestIp"></el-input>
                    </el-form-item>
                    <el-form-item label="地址" prop="address">
                        <el-input v-model="m.address"></el-input>
                    </el-form-item>
                    <el-form-item label="描述" prop="description">
                        <el-input v-model="m.description"></el-input>
                    </el-form-item>
                    <el-form-item label="浏览器" prop="browser" min-width="100" :show-overflow-tooltip="true">
                        <el-input v-model="m.browser"></el-input>
                    </el-form-item>
                    <el-form-item label="请求耗时" prop="time">
                        <el-input v-model="m.time"></el-input>
                    </el-form-item>
                    <el-form-item label="请求方式" prop="method">
                        <el-input v-model="m.method"></el-input>
                    </el-form-item>
                    <el-form-item label="参数" prop="params" min-width="100" :show-overflow-tooltip="true">
                        <el-input v-model="m.params"></el-input>
                    </el-form-item>
                    <el-form-item label="日志类型" prop="logType">
                        <el-input v-model="m.logType"></el-input>
                    </el-form-item>
                    <el-form-item label="异常详情" prop="exceptionDetail">
                        <el-input v-model="m.exceptionDetail"></el-input>
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
//import inputEnum from "../../sa-resources/com-view/input-enum.vue";
export default {
    //components: { inputEnum },
    props: ["params"],
    data() {
        return {
            m: {},
            title: "",
            isShow: false,
            rules: {
                requestIp: [],
                address: [],
                description: [
                    {
                        min: 0,
                        max: 2000,
                        message: "长度在 2000 个字符",
                        trigger: "blur",
                    },
                ],
                browser: [],
                time: [],
                method: [],
                params: [
                    {
                        min: 0,
                        max: 2000,
                        message: "长度在 2000 个字符",
                        trigger: "blur",
                    },
                ],
                logType: [],
                exceptionDetail: [],
            },
            fileList: [],
        };
    },
    methods: {
        open: function (data) {
            this.isShow = true;
            if (data) {
                this.title = "修改 系统日志";

                this.m = data;
            } else {
                this.m = {
                    requestIp: "",
                    address: "",
                    description: "",
                    browser: "",
                    method: "",
                    params: "",
                    logType: "",
                    exceptionDetail: "",
                };
                this.title = "添加 系统日志";
            }
        },

        //提交系统日志信息
        ok: function (formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.sa.post("/Log/save", this.m).then((res) => {
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
    created() {},
};
</script>