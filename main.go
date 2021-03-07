package main

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"regexp"
	"strings"
	"time"
	"yixiang.co/yshop/common"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/jwt"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/redis"
	_ "yixiang.co/yshop/routers"
	"yixiang.co/yshop/untils"
)


const bearerLength = len("Bearer ")

func init()  {
	//初始化数据库
	dbUser, _ := beego.AppConfig.String("mysqluser")
	dbPass, _ := beego.AppConfig.String("mysqlpass")
	dbName, _ := beego.AppConfig.String("mysqldb")
	dbHost, _ := beego.AppConfig.String("mysqlhost")
	dbPort, _ := beego.AppConfig.String("mysqlport")
	maxIdleConn, _ := beego.AppConfig.Int("db_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("db_max_open_open")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",
		dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=true&loc=Asia%2FShanghai",
		)
	orm.MaxIdleConnections(maxIdleConn)
	orm.MaxOpenConnections(maxOpenConn)
	orm.DefaultTimeLoc = time.UTC

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}

	//日志
	logs.Async()
	level, _ := beego.AppConfig.Int("logLevel")
	logs.SetLevel(level)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./logs/yshop.log",
	"level":6,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30,
	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info"]}`)

}

func main() {
	log.Print("==========yshop管理后台 start=============")
	//if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//


	beego.InsertFilterChain("/*", func(next beego.FilterFunc) beego.FilterFunc {
		return func(ctx *context.Context) {
			url := ctx.Input.URL()
			ignoreUrls, _ := beego.AppConfig.String("ignore_urls")
			if strings.Contains(ignoreUrls,url) || strings.Contains(url,"/swagger")  || strings.Contains(url,"/static") {
					next(ctx)
				}else {
					method := strings.ToLower(ctx.Input.Method())
					//部署线上开启
					prohibit := "post,put,delete"
					if url != "/admin/auth/logout" && strings.Contains(prohibit,method) {
						ctx.Output.JSON(controllers.ErrMsg("演示环境禁止操作",40006),
							true,true)
						return
					}
					mytoken := ctx.Input.Header("Authorization")
					if len(mytoken) < bearerLength {
						ctx.Output.Status = 401
						ctx.Output.JSON(controllers.ErrMsg("header Authorization has not Bearer token",40001),
							true,true)
						return
					}
					token := strings.TrimSpace(mytoken[bearerLength:])
					usr, err := jwt.ValidateToken(token)
					if err != nil {
						ctx.Output.Status = 401
						ctx.Output.JSON(controllers.ErrMsg(err.Error(),40001),
							true,true)
						return
					}
					//校验权限
					index := checkPermission(url,method,token)
					if index == -1 {
						ctx.Output.JSON(controllers.ErrMsg("无权限",40001),true,true)
						return
					}

					ctx.Input.SetData(common.ContextKeyUserObj,usr)
					next(ctx)
			}
		}
	})
	beego.Run()

}

func checkPermission(url string, method string, token string) int  {
	logs.Info(url,method,token,method)
	//公共路由直接放行
	var ignoreUrls = "/admin/menu/build,/admin/user/center,/admin/user/updatePass,/admin/auth/info,/admin/auth/logout"
	if strings.Contains(ignoreUrls,url) {
		return 0
	}
	reg := regexp.MustCompile(`[0-9]+`)
	newUrl := reg.ReplaceAllString(url,"*")
	permission := models.FindByRouterAndMethod(newUrl,method)
	var key = common.REDIS_PREFIX_AUTH + token
	userMap, _:= redis.Get(key)
	jsonStr := userMap[key]
	user := &models.User{}
	json.Unmarshal([]byte(jsonStr),user)

	index := untils.Contains(user.Permissions,permission)

	return index
}
