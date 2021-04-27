package models

import "github.com/beego/beego/v2/client/orm"

type SysLog struct {
	Id     int64
	Description string
	ExceptionDetail string
	LogType string
	Method string
	Params string
	RequestIp string
	Username string
	Address string
	Browser string
	Type int8
	Uid int64
	BaseModel
}

func init() {
	orm.RegisterModel(new(SysLog))
}