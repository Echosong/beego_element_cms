/* 
	在开发时，如无必要，请不要直接魔改模板的代码，以免在运行时出现意外bug
	这是模板为你预留一个文件，用来对接你的业务逻辑 
 */

// import Vue from 'vue';
export default function(sa_admin, sa) {


	// ================================= 示例：模板基础信息  =================================
	sa_admin.title = '上海开赟软件';
	// sa_admin.logo_url = '图片地址';    // 设置logo图标地址   默认值空, 代表使用：./admin-logo.png
	// sa_admin.icon_url = '图片地址';    // 设置icon图标地址   默认值空, 代表使用：./favicon.ico


	// ================================= 示例：初始化菜单  =================================
	sa_admin.initMenu(); // 初始化菜单, 不传参代表默认显示所有菜单 菜单在 ./sa-menu-list.js 里,
	// sa_admin.initMenu(['1', '1-1', '1-2', '4']); 	// 传入一个id数组, 显示指定菜单  
 
 
	// ================================= 示例：设置头像昵称  =================================
	// 
	sa_admin.$nextTick(function() {
		
		this.sa.get('/user/getCurrent').then(res => {
			let user = res.data;
			this.sa_admin.user = user;
		})
		document.title = sa_admin.title ;
	})
	
	
	// ================================= 示例：js操作菜单  =================================
	
	// sa_admin.showHome();            // 显示主页选项卡 
	// sa_admin.showTabById('home');    // 显示一个选项卡, 根据id
	// sa_admin.closeTabById('5-1');    // 关闭一个选项卡，根据 id （ 第二个参数可填关闭后的回调函数 ）
	// sa_admin.showMenuById('5-2');    // 打开一个 菜单，根据 id
	
	// 跨窗口通信
	// sa_admin.getView('5-2').f5();// 获取指定视图的组件对象，并调用其f5()函数，一般用于跨窗口通信 
	
	// 以下写法，调用tab打开新窗口
	// var tab = {id: '5-1', name: '用户添加', view: () => import('@/sa-view/user/user-add.vue'), params: {username: '王铁汉2'} };
	// sa_admin.showTab(tab);
	

	// ================================= 示例：设置登录后的头像处，下拉可以出现的选项  =================================
	sa_admin.dropList = [ // 头像点击处可操作的选项
		{
			name: '重置密码',
			click: function() {
				sa_admin.userPassword();
			}
		},
		{
			name: '切换账号',
			click: function() {
				sa_admin.openLogin();	// 打开登陆视图 
			}
		},
		{
			name: '退出登录',
			click: function() {
				// sa_admin.$message('点击了退出登录，你可以参照文档重写此函数');
				sa.confirm('退出登录？', function() {
					sa.get('/user/logout').then(res=>{
						localStorage.removeItem("token");
						sa.alert('注销成功', function() {
							sa_admin.openLogin();
						})
					})
				});
			}
		}
	]


	// 初始化模板, 必须调用
	sa_admin.init();


}
