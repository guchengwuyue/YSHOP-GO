package initialize

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

func dbinit()  {
	runmode,_ := beego.AppConfig.String("runmode")
	isDev := (runmode == "dev")
	registDatabase()
	if isDev {
		orm.Debug = isDev
	}
}

func registDatabase()  {
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

}
