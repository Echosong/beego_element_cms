package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//用户
	orm.RegisterModelWithPrefix("hz_", new(User))
	//角色
	orm.RegisterModelWithPrefix("hz_", new(Role))
	//用户角色关系
	orm.RegisterModelWithPrefix("hz_", new(UserRole))
	//权限
	orm.RegisterModelWithPrefix("hz_", new(Permission))
	//权限角色关系
	orm.RegisterModelWithPrefix("hz_", new(RolePermission))
	//用户操作日志
	orm.RegisterModelWithPrefix("hz_", new(UserLog))
	//文章分类
	orm.RegisterModelWithPrefix("hz_", new(Category))
	//配置信息
	orm.RegisterModelWithPrefix("hz_", new(Config))
	//内容
	orm.RegisterModelWithPrefix("hz_", new(Content))
	//广告
	orm.RegisterModelWithPrefix("hz_", new(Ad))
	//产品
	orm.RegisterModelWithPrefix("hz_", new(Product))
	//文章管理
	orm.RegisterModelWithPrefix("hz_", new(Article))

	dbMap, _ := config.GetSection("mysql")
	orm.RegisterDataBase("default", "mysql", dbMap["conn"]+"?charset=utf8")
	// automatically build table
	// 开发模式才处理数据自动添加问题
	if beego.BConfig.RunMode == "dev" {
		orm.RunSyncdb("default", false, true)
	}
}
