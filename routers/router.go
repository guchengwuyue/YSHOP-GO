// @APIVersion 1.0.0
// @Title YSHOP-GO API
// @Description YSHOP-GO管理系统API
// @TermsOfServiceUrl https://www.yixiang.co/
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"yixiang.co/yshop/controllers/admin"
)

func init() {

	//移动端路由
	ns1 := beego.NewNamespace("/v1",
		beego.NSNamespace("/mem",
			beego.NSInclude(

			),
		),
	)
	//管理后台路由
	ns2 := beego.NewNamespace("/admin",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&admin.LoginController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&admin.UserController{},
			),
		),
		beego.NSNamespace("/menu",
			beego.NSInclude(
				&admin.MenuController{},
			),
		),
		beego.NSNamespace("/dict",
			beego.NSInclude(
				&admin.DictController{},
			),
		),
		beego.NSNamespace("/dictDetail",
			beego.NSInclude(
				&admin.DictDetailController{},
			),
		),
		beego.NSNamespace("/dept",
			beego.NSInclude(
				&admin.DeptController{},
			),
		),
		beego.NSNamespace("/job",
			beego.NSInclude(
				&admin.JobController{},
			),
		),
		beego.NSNamespace("/roles",
			beego.NSInclude(
				&admin.RoleController{},
			),
		),
		beego.NSNamespace("/logs",
			beego.NSInclude(
				&admin.LogController{},
			),
		),
	)
	beego.AddNamespace(ns1,ns2)




	//beego.SetStaticPath("/swagger/", "swagger")
}
