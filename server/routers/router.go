package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"kaiyu-web/controllers"
	"kaiyu-web/controllers/admin"
)

func init() {

	//前端模板
	beego.Include(&controllers.MainController{})

	//后端接口
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/main",
			beego.NSInclude(
				&admin.MainController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&admin.UserController{},
			),
		),
		beego.NSNamespace("/role",
			beego.NSInclude(
				&admin.RoleController{},
			),
		),
		beego.NSNamespace("/config",
			beego.NSInclude(
				&admin.ConfigController{},
			),
		),
		beego.NSNamespace("/ad",
			beego.NSInclude(
				&admin.AdController{},
			),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(
				&admin.ArticleController{},
			),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(
				&admin.CategoryController{},
			),
		),
		beego.NSNamespace("/permission",
			beego.NSInclude(
				&admin.PermissionController{},
			),
		),
		beego.NSNamespace("/content",
			beego.NSInclude(
				&admin.ContentController{},
			),
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&admin.ProductController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
