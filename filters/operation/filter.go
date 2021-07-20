package operation

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
	"time"
	"yixiang.co/yshop/common/jwt"
	"yixiang.co/yshop/models"
	_ "yixiang.co/yshop/routers"
)



type FilterChainBuilder struct {

}
//记录操作日志
func (builder *FilterChainBuilder) AccessLog(next beego.FilterFunc) beego.FilterFunc {
	return func(ctx *context.Context) {
		url := ctx.Input.URL()
		user, err := jwt.GetAdminUser(ctx.Input)
		if err != nil {
			next(ctx)
			return
		}
		method := strings.ToLower(ctx.Input.Method())
		reg := regexp.MustCompile(`[0-9]+`)
		newUrl := reg.ReplaceAllString(url,"*")
		menu := models.FindMenuByRouterAndMethod(newUrl,method)
		log := models.SysLog{
			Description:     menu.Name,
			Method:          method,
			RequestIp:       ctx.Input.IP(),
			Username:        user.Username,
			Address:         newUrl,
			Browser:         "",
			Type:            0,
			Uid:             user.Id,
		}
		now := time.Now()
		next(ctx)
		consume := time.Now().Sub(now)
		log.Time = consume.Microseconds()
		models.AddLog(&log)


	}
}


