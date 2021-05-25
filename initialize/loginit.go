package initialize

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func loginit()  {
	//日志
	logs.Async()
	level, _ := beego.AppConfig.Int("logLevel")
	logs.SetLevel(level)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./logs/yshop.log",
	"level":6,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30,
	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info"]}`)
}
