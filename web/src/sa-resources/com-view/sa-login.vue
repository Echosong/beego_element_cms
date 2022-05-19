<style scoped>
/* 视图盒子 */
.view-box {
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0px;
    left: 0px;
    z-index: 1000;
}
/* EAEFF3 */
.bg-1 {
    height: 50%;
    background: linear-gradient(to bottom right, #0466c5, #3496f5);
}
.bg-2 {
    height: 50%;
    background-color: #eaeff3;
}

.login-box {
    width: 400px;
    height: 400px;
    position: absolute;
    left: calc(50% - 200px);
    top: calc(50% - 250px);
}
/* .login-box{} */

/* logo */
.login-top {
    margin-bottom: 40px;
    text-align: center;
}
.logo-img {
    width: 50px;
    height: 50px;
    vertical-align: middle;
    margin-left: -10px;
    margin-right: 20px;
}
.logo-img {
    position: relative;
    top: -5px;
}
.admin-title {
    font-size: 28px;
    color: #fff;
    font-weight: 700;
}

/* 表单 */
.from-box {
    padding: 20px 50px;
    background-color: #fff;
}
.from-box {
    border-radius: 1px;
    box-shadow: 1px 1px 2px #666;
}
.from-title {
    margin-top: 20px;
    margin-bottom: 30px;
    text-align: center;
}
</style>
<template>
    <div class="view-box" v-if="isShow">
        <div class="bg-1"></div>
        <div class="bg-2"></div>
        <div class="login-box">
            <div class="login-box-2">
                <div class="login-top">
                    <img src="./../index/logo.png" class="logo-img">
                    <span class="admin-title">上海开赟软件</span>
                </div>
                <div class="from-box">
                    <h3 class="from-title">账号登录</h3>
                    <el-form size="small" label-position="left" label-width="0px">
                        <el-form-item>
                            <el-input prefix-icon="el-icon-user" v-model="m.username" placeholder="请输入账号"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-input prefix-icon="el-icon-unlock" v-model="m.password" type="password"
                                placeholder="请输入密码" @keyup.native.enter="ok()"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" size="small" style="width: 100%;" @click="ok()">登录</el-button>
                        </el-form-item>
                        <el-form-item>
                            <span style="color: #999;"></span>
                            <el-link style="float: right;" @click="isShow=false">注册</el-link>
                        </el-form-item>
                    </el-form>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
let Base64 = require('js-base64').Base64;
export default {
    name: 'sa-login',
    data() {
        return {
            isShow: false,	// 是否显示当前视图 
            m: {
                username: '',
                password: ''
            },
        }
    },
    methods: {
        // 点击确定 
        ok() {
            // 表单验证  
            if (this.m.username == '' || this.m.password == '') {
                return this.sa.error2('请输入完整');
            }
            this.m.password = Base64.encode(this.m.password);
            this.sa.post('/user/login', this.m).then(res => {
				console.log("信息登录后信息==="+res);
				
				localStorage.setItem("token", res.data);
                console.log(res);
                location.href = "/";
                
            })
            this.m.password = '';
        }
    }
}
</script>

