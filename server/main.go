package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"kaiyu-web/auto"
	_ "kaiyu-web/filters"
	_ "kaiyu-web/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		//自动生成视图
		auto.NewView()
	}
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}
