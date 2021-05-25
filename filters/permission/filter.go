package permission

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
	"yixiang.co/yshop/common"
	"yixiang.co/yshop/common/jwt"
	"yixiang.co/yshop/common/redis"
	"yixiang.co/yshop/common/untils"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/models"
	_ "yixiang.co/yshop/routers"
)

const bearerLength = len("Bearer ")

type FilterChainBuilder struct {

}

func (builder *FilterChainBuilder) Permssion(next beego.FilterFunc) beego.FilterFunc {
	return func(ctx *context.Context) {
		url := ctx.Input.URL()
		ignoreUrls, _ := beego.AppConfig.String("ignore_urls")
		if strings.Contains(ignoreUrls,url) || strings.Contains(url,"/swagger")  || strings.Contains(url,"/static") {
			next(ctx)
		}else {
			method := strings.ToLower(ctx.Input.Method())
			//部署线上开启
			//prohibit := "post,put,delete"
			//if url != "/admin/auth/logout" && strings.Contains(prohibit,method) {
			//	ctx.Output.JSON(controllers.ErrMsg("演示环境禁止操作",40006),
			//		true,true)
			//	return
			//}
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
	user := &models.SysUser{}
	json.Unmarshal([]byte(jsonStr),user)

	index := untils.Contains(user.Permissions,permission)

	return index
}
