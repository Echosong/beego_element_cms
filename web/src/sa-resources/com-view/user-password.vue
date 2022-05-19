<template>
    <el-dialog v-if="m" title="修改密码" :visible.sync="isShow" width="550px" top="10vh" :append-to-body="true"
        :destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
        <div class="vue-box">
            <div class="c-panel">
                <el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm" label-width="120px">
                    <el-form-item label="原始密码" prop="password">
                        <el-input placeholder="请输入原始密码" v-model="m.password" show-password></el-input>
                    </el-form-item>
                    <el-form-item label="新密码" prop="newPassword">
                        <el-input placeholder="请输入新密码" v-model="m.newPassword" show-password></el-input>
                    </el-form-item>
                    <el-form-item label="确认密码" prop="rePassword">
                        <el-input placeholder="请输入确认密码" v-model="m.rePassword" show-password></el-input>
                    </el-form-item>

                    <el-form-item>
                        <span class="c-label">&emsp;</span>
                        <el-button type="primary" icon="el-icon-setting" size="mini" @click="ok('ruleForm')">确定
                        </el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </el-dialog>
</template>

<script>
export default {
    data() {
        return {
            m: {
                password: '',
                newPassword: '',
                rePassword: '',
            },
            rules: {
                password: [{ required: true, message: '请输入原始密码', trigger: 'blur' },
                { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }],
                newPassword: [{ required: true, message: '请输入新密码', trigger: 'blur' },
                { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }],
            },
            isShow: false
        }
    },
    methods: {
        ok: function (formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.sa.post("/user/setPassword", this.m).then(res => {
                        this.sa.ok("密码修改成功");
                        this.isShow = false;
                    })
                }
            })
        }
    }
}
</script>

<style>
</style>