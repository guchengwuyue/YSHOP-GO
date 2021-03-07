package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"yixiang.co/yshop/dto"
)

type Dict struct {
	Id     int64 `json:"id"`
	Name string `json:"name" valid:"Required;"`
	Remark string `json:"remark" valid:"Required;"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Dict))
}

// get all
func GetAllDict(base dto.BasePage,query ...interface{}) (int,[]Dict)  {
	var (
		tableName = "dict"
		dicts []Dict
		condition = ""
	)
	if base.Blurry != "" {
		condition = " and name= '" + base.Blurry + "'"
	}
	logs.Info(base)
	total, _, rs := GetPagesInfo(tableName, base.Page, base.Size, condition)
	rs.QueryRows(&dicts)

	//o := orm.NewOrm()

	return total,dicts
}

// last inserted Id on success.
func AddDict(m *Dict) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByDict(m *Dict) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByDict(id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE dict SET is_del = ? WHERE id = ?", 1, id).Exec()
	return
}