package models

import (
	"github.com/beego/beego/v2/client/orm"
	"yixiang.co/yshop/common/untils"
	"yixiang.co/yshop/models/dto"
)

type SysLog struct {
	Id     int64 `json:"id"`
	Description string `json:"description"`
	Method string `json:"method"`
	RequestIp string `json:"requestIp"`
	Username string `json:"username"`
	Address string `json:"address"`
	Browser string `json:"browser"`
	Type int8 `json:"type"`
	Uid int64 `json:"uid"`
	Time int64 `json:"time"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(SysLog))
}

// get all
func GetAllLog(base dto.BasePage,query ...interface{}) (int,[]SysLog)  {
	var (
		tableName = "sys_log"
		lists []SysLog
		condition = ""
	)
	if base.Blurry != "" {
		condition = " and username= '" + base.Blurry + "'"
	}


	total, _, rs := GetPagesInfo(tableName, base.Page, base.Size, condition)
	rs.QueryRows(&lists)

	return total,lists
}


// last inserted Id on success.
func AddLog(m *SysLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func DelBylog(ids []int64) (err error) {
	str := untils.ReturnQ(len(ids))
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE sys_log SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
	return
}
