<style scoped>
</style>

<template>
	<div class="vue-box">
		<!-- 参数栏 -->
		<div class="c-panel">
			<!-- 参数栏 -->
		
			<!-- 数据列表 -->
			<el-table :data="dataList" size="mini" >
				<el-table-column label="编号" prop="id" width="70px" > </el-table-column>
				<el-table-column label="名称">
					<template slot-scope="s">
						<el-input size="mini" v-model="s.row.name" @input="s.row.is_update = true"></el-input>
					</template>
				</el-table-column>
				<el-table-column label="描述">
					<template slot-scope="s">
						<el-input size="mini" v-model="s.row.description" @input="s.row.is_update = true"></el-input>
					</template>
				</el-table-column>
				<el-table-column label="创建日期">
					<template slot-scope="s">
						{{s.row.createTime}}
					</template>
				</el-table-column>
				<el-table-column label="操作" width="220px">
					<template slot-scope="s">
						<el-button type="text" size="mini" @click="update(s.row)">
							<span :style="s.row.is_update ? 'color: red;' : ''">修改</span>
						</el-button>
						<el-button type="text" size="mini" @click="del(s.row)">删除</el-button>
						<el-button type="text" size="mini" @click="menu_setup(s.row)">分配权限</el-button>
					</template>
				</el-table-column>
			</el-table>
		</div>
		<!-- 添加 -->
		<div class="c-panel">
			<h4 class="c-title">添加一个</h4>
			<el-table :data="[{}]" size="mini" >
				<el-table-column label="名称">
					<template>
						<el-input size="mini" v-model="m.name"></el-input>
					</template>
				</el-table-column>
				<el-table-column label="描述">
					<template>
						<el-input size="mini" v-model="m.description"></el-input>
					</template>
				</el-table-column>
				<el-table-column prop="address" label="操作">
					<template>
						<el-button type="text" size="mini" @click="add">添加</el-button>
					</template>
				</el-table-column>
			</el-table>
		</div>
		
		<!-- 设置权限  -->
		<menu-setup ref="menu-setup"></menu-setup>
		
		
		
	</div>
</template>

<script>
	
	//import mockData from './mock-data.js';	 	// 模拟数据
	import menuSetup from './menu-setup.vue';	 	// 设置权限 
	export default {
		components: {
			menuSetup
		},
		data() {
			return {
				isShow: true,
				p: {	// 查询参数 
					name: '',
				},
				dataList: [],	// 数据集合
				m: {		// 添加信息
					id: 0, 
					name: '部门名称',
					description: '描述',
					is_lock: 2,
					is_update: false,
				},
				rid: 0
			}
		},
		methods: {
			// 刷新
			f5: function(){
				this.sa.get("/role/list").then(res=>{
					console.log(res);
					this.dataList = res.data;
				});
			},
			// 修改
			update: function (data) {
				this.sa.post('/role/save', data).then(res=>{
					console.log(res);
					this.sa.ok("修改成功")
					data.is_update = false;
				})
			},
			// 删除
			del: function (data) {
				this.sa.confirm('是否删除，此操作不可撤销', function(){
					this.sa.delete('/role/delete/'+data.id).then(res=>{
						console.log(res);
						this.sa.ok("删除成功");
						this.f5();
					})
				}.bind(this))
			},
			// 添加
			add: function () {
				var data = this.sa.copyJSON(this.m);
				this.sa.post('/role/save', data).then(res=>{
					console.log(res);
					this.sa.ok("添加成功");
					this.f5();
				});
			}, 
			// 修改菜单权限 
			menu_setup: function(data){
				this.$refs['menu-setup'].show(data.id, data.name);
			}
		},
		created: function(){
			this.f5();
		}
	}
</script>
