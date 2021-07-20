package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"yixiang.co/yshop/filters/operation"
	"yixiang.co/yshop/filters/permission"
	_ "yixiang.co/yshop/initialize"
	_ "yixiang.co/yshop/routers"
)


func main() {
	log.Print("==========yshop管理后台 start=============")
	//if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//

	//人员操作日志filter
	operation := &operation.FilterChainBuilder{}
	beego.InsertFilterChain("/*", operation.AccessLog)

	//权限filter
	fb := &permission.FilterChainBuilder{}
	beego.InsertFilterChain("/*", fb.Permssion)




	beego.Run()

}


