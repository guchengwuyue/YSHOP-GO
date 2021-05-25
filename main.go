package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"yixiang.co/yshop/filters/permission"
	_ "yixiang.co/yshop/routers"
	_ "yixiang.co/yshop/initialize"
)


func main() {
	log.Print("==========yshop管理后台 start=============")
	//if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//
	fb := &permission.FilterChainBuilder{}
	beego.InsertFilterChain("/*", fb.Permssion)


	beego.Run()

}


