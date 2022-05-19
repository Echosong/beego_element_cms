package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"],
		beego.ControllerComments{
			Method:           "ListPage",
			Router:           `/listPage`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:AdController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/find/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"],
		beego.ControllerComments{
			Method:           "ListPage",
			Router:           `/listPage`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ArticleController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/find/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"],
		beego.ControllerComments{
			Method:           "ListPage",
			Router:           `/listPage`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:CategoryController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ConfigController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ConfigController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/find/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ConfigController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ConfigController"],
		beego.ControllerComments{
			Method:           "All",
			Router:           `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ConfigController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ConfigController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/:code`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"],
		beego.ControllerComments{
			Method:           "ListPage",
			Router:           `/listPage`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ContentController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:MainController"],
		beego.ControllerComments{
			Method:           "Upload",
			Router:           `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/find/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method:           "ListByRole",
			Router:           `/listByRole/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method:           "ListByUser",
			Router:           `/listByUser/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method:           "ListPage",
			Router:           `/listPage`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method:           "UpdateRolePermissions",
			Router:           `/updateRolePermissions/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/list/:ids`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:ProductController"],
		beego.ControllerComments{
			Method:           "SaveList",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"DELETE"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/find/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method:           "ListPage",
			Router:           `/listPage`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "Find",
			Router:           `/find/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "GetCurrent",
			Router:           `/getCurrent`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "ListPage",
			Router:           `/listPage`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "SetPassword",
			Router:           `/setPassword`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "About",
			Router:           `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Case",
			Router:           `/case`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Contact",
			Router:           `/contact`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Partner",
			Router:           `/partner`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Product",
			Router:           `/product`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Quality",
			Router:           `/quality`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Search",
			Router:           `/search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Service",
			Router:           `/service`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"] = append(beego.GlobalControllerRouter["kaiyu-web/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Content",
			Router:           `content/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
