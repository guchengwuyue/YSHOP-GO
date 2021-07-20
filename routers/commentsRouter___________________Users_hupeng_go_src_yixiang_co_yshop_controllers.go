package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DeptController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:DictDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:JobController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "Captcha",
            Router: "/captcha",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "Info",
            Router: "/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Build",
            Router: "/build",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "GetTree",
            Router: "/tree",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Menu",
            Router: "/menu",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:ToolsController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:ToolsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Center",
            Router: "/center",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Avatar",
            Router: "/updateAvatar",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["yixiang.co/yshop/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Pass",
            Router: "/updatePass",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
