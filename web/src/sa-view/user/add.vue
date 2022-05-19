<template>
	<el-dialog v-if="m" :title="title" :visible.sync="isShow" width="550px" top="10vh" :append-to-body="true"
		:destroy-on-close="true" :close-on-click-modal="false" custom-class="full-dialog">
		<div class="vue-box">
			<!-- 参数栏 -->
			<div class="c-panel">
				<el-form size="mini" v-if="m" ref="ruleForm" :rules="rules" :model="m" class="demo-ruleForm"
					label-width="120px">
					<el-form-item label="用户名" prop="username">
						<el-input v-model="m.username" :disabled="m.id"></el-input>
					</el-form-item>
					<el-form-item label="姓名" prop="name">
						<el-input v-model="m.name"></el-input>
					</el-form-item>
					<el-form-item label="邮件" prop="email">
						<el-input v-model="m.email"></el-input>
					</el-form-item>
					<el-form-item label="性别：">
						<el-radio-group v-model="m.sex" size="mini">
							<el-radio :label="0">男</el-radio>
							<el-radio :label="1">女</el-radio>
						</el-radio-group>
					</el-form-item>
					<el-form-item label="密码" prop="password">
						<el-input v-model="m.password" show-password></el-input>
					</el-form-item>
					<el-form-item label="部门" prop="role_id">
						<select-role v-model="m.role_id" ref="Role"  ></select-role>
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
	import selectRole from "../../sa-resources/com-view/select-role.vue";
	export default {
		components: {
			selectRole
		},
		props: ["params"],
		data() {
			return {
				m: {username:''},
				title: "",
				isShow: false,
				rules: {
					username: [{
						required: true,
						message: '请输入用户名',
						trigger: 'blur'
					}, ],
					name: [],
					state: [],
					email: [{
						type: 'email',
						message: '请输入正确的邮箱地址',
						trigger: ['blur', 'change']
					}, ],
					regIp: [],
					loginIp: [],
					password: [],
					roleId: [],
				},
				fileList: [],
			}
		},
		methods: {
			open: function(data) {
				this.isShow = true;
				if (data) {
					this.title = "修改用户表";
					this.m = data;
				} else {
					this.m = {
						username: "",
						name: "",
						sex: 1,
						email: '',
						password: "",
						role_id: 1
					}
					this.title = "添加 用户表";
				}
	 			this.$refs['Role'].setValue(this.m.role_id)

			},


			//提交用户表信息
			ok: function(formName) {
				this.$refs[formName].validate((valid) => {
					if (valid) {

						this.sa.post("/user/save", this.m).then((res) => {
							console.log(res);
							this.$parent.f5();
							this.isShow = false;
						});
					} else {
						console.log("error submit!!");
						return false;
					}
				});
			}
		},
		created() {},
	};
</script>
