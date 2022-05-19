<style scoped>
</style>

<template>
    <el-dialog v-if="m" :title="title" :visible.sync="isShow" width="500px" top="10vh" :append-to-body="true" :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <!-- 参数栏 -->
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm" label-width="120px">

                    <el-form-item label="姓名：" prop="name" required>
                        <el-input v-model="m.name"></el-input>
                    </el-form-item>

                    <el-form-item label="手机号：" v-if="!m.id" prop="username" required>
                        <el-input v-model="m.username"></el-input>
                    </el-form-item>

                    <el-form-item label="密码：" v-if="!m.id" prop="password" required>
                        <el-input type="password" v-model="m.password"></el-input>
                    </el-form-item>

                    <el-form-item label="邮箱：">
                        <el-input v-model="m.email"></el-input>
                    </el-form-item>

                    <el-form-item label="部门：">
                        <el-select v-model="m.roleId">
                            <el-option label="请选择" value="0" disabled></el-option>
                            <el-option v-for="item in roles" :key="item.id" :label="item.name" :value="item.id">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="性别：">
                        <el-radio-group v-model="m.sex" size="mini">
                            <el-radio :label="0">男</el-radio>
                            <el-radio :label="1">女</el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item label="状态">
                        <el-tag :type="m.state ==2 ?'danger':'' ">{{m.stateEnum}}</el-tag>
                    </el-form-item>
                    <el-form-item>
                        <span class="c-label">&emsp;</span>
                        <el-button type="primary" icon="el-icon-plus" size="mini" @click="ok('ruleForm')">确定</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </el-dialog>
</template>

<script>
export default {
    props: ["params"],
    data() {
        return {
            m: {
                username: "", // 从菜单配置文件里传递过来的参数
                password: "",
                name: "",
                email: "",
                roleId: 0,
                sex: 1,
                state: 0
            },
            rules: {
                name: [
                    {
                        required: true,
                        message: "请输入姓名",
                        trigger: "blur"
                    }
                ],
                username: [
                    { required: true, message: "请输入账号", trigger: "blur" },
                    {
                        min: 10,
                        max: 12,
                        message: "账号必须为手机号",
                        trigger: "blur"
                    }
                ],
                password: [
                    { required: true, message: "请输入密码", trigger: "blur" },
                    { main: 6, max: 24, message: "密码长度必须大于等于6个字符" }
                ]
            },
            isShow: false,
            title: "新增用户",
            roles: [{ name: "管理员", id: 1 }]
        };
    },
    methods: {
        open: function(data) {
            this.init();
            this.isShow = true;
            if (data.id) {
                data.password = "";
                this.title = "修改用户";
                this.m = data;
            
            }
        },

        //提交用户信息
        ok: function(formName) {
            this.$refs[formName].validate(valid => {
                if (valid) {
                    if (this.m.id) {
                        this.sa.put("/user/update", this.m).then(res => {
                            console.log(res);
                            this.$parent.f5();
                            this.isShow = false;
                        });
                    } else {
                        this.sa.post("/user/create", this.m).then(res => {
                            console.log(res);
                            this.$parent.f5();
                            this.isShow = false;
                        });
                    }
                } else {
                    console.log("error submit!!");
                    return false;
                }
            });
        },
        /**
         * 获取角色
         */
        init() {
            this.sa.get("/role/list").then(res => {
                console.log(res);
                this.roles = res.data;
            });
        }
    },
    created() {}
};
</script>


